// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: internal/pkg/fetch/git/git.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package git

import (
	"fmt"
	"regexp"

	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/bomctl/bomctl/internal/pkg/url"
	"github.com/bomctl/bomctl/internal/pkg/utils"
)

type GitFetcher struct{}

func (gf *GitFetcher) RegExp() *regexp.Regexp {
	return regexp.MustCompile(
		fmt.Sprintf("%s%s%s%s%s",
			`(?P<scheme>git|ssh)(?:@|(\+https?)?://)`,
			`((?P<username>[^:]+)(?::(?P<password>[^@]+))?(?:@))?`,
			`(?P<hostname>[^@/?#:]*)(?::(?P<port>\d+)?)?`,
			`(?P<path>[^@?#]*)(?:@(?P<gitRef>[^#]+))?`,
			`(\?(?P<query>[^#]*))?(#(?P<fragment>.*))?`,
		),
	)
}

func (gf *GitFetcher) Parse(fetchURL string) *url.ParsedURL {
	results := map[string]string{}
	pattern := gf.RegExp()
	match := pattern.FindStringSubmatch(fetchURL)

	for idx, name := range match {
		results[pattern.SubexpNames()[idx]] = name
	}

	return &url.ParsedURL{
		Scheme:   results["scheme"],
		Username: results["username"],
		Password: results["password"],
		Hostname: results["hostname"],
		Port:     results["port"],
		Path:     results["path"],
		GitRef:   results["gitRef"],
		Query:    results["query"],
		Fragment: results["fragment"],
	}
}

func (gf *GitFetcher) Fetch(parsedURL *url.ParsedURL, auth *url.BasicAuth) (*sbom.Document, error) {
	memStorage := memory.NewStorage()
	memFS := memfs.New()

	refName := plumbing.NewRemoteReferenceName("origin", parsedURL.GitRef)

	repository, err := git.Clone(memStorage, memFS, &git.CloneOptions{
		URL:           parsedURL.String(),
		Auth:          auth,
		RemoteName:    "origin",
		ReferenceName: refName,
		SingleBranch:  true,
		Depth:         1,
		ProxyOptions:  transport.ProxyOptions{},
	})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	tree, err := repository.Worktree()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	sbomFile, err := tree.Filesystem.Open(parsedURL.Fragment)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	sbomBytes := []byte{}
	_, err = sbomFile.Read(sbomBytes)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	document, err := utils.ParseSBOMData(sbomBytes)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return document, nil
}
