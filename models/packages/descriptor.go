// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package packages

import (
	"context"
	"fmt"
	"net/url"

	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/json"
	"code.gitea.io/gitea/modules/packages/composer"
	"code.gitea.io/gitea/modules/packages/conan"
	"code.gitea.io/gitea/modules/packages/container"
	"code.gitea.io/gitea/modules/packages/helm"
	"code.gitea.io/gitea/modules/packages/maven"
	"code.gitea.io/gitea/modules/packages/npm"
	"code.gitea.io/gitea/modules/packages/nuget"
	"code.gitea.io/gitea/modules/packages/pypi"
	"code.gitea.io/gitea/modules/packages/rubygems"

	"github.com/hashicorp/go-version"
)

// PackagePropertyList is a list of package properties
type PackagePropertyList []*PackageProperty

// GetByName gets the first property value with the specific name
func (l PackagePropertyList) GetByName(name string) string {
	for _, pp := range l {
		if pp.Name == name {
			return pp.Value
		}
	}
	return ""
}

// PackageDescriptor describes a package
type PackageDescriptor struct {
	Package           *Package
	Owner             *user_model.User
	Repository        *repo_model.Repository
	Version           *PackageVersion
	SemVer            *version.Version
	Creator           *user_model.User
	PackageProperties PackagePropertyList
	VersionProperties PackagePropertyList
	Metadata          interface{}
	Files             []*PackageFileDescriptor
}

// PackageFileDescriptor describes a package file
type PackageFileDescriptor struct {
	File       *PackageFile
	Blob       *PackageBlob
	Properties PackagePropertyList
}

// PackageWebLink returns the package web link
func (pd *PackageDescriptor) PackageWebLink() string {
	return fmt.Sprintf("%s/-/packages/%s/%s", pd.Owner.HTMLURL(), string(pd.Package.Type), url.PathEscape(pd.Package.LowerName))
}

// FullWebLink returns the package version web link
func (pd *PackageDescriptor) FullWebLink() string {
	return fmt.Sprintf("%s/%s", pd.PackageWebLink(), url.PathEscape(pd.Version.LowerVersion))
}

// CalculateBlobSize returns the total blobs size in bytes
func (pd *PackageDescriptor) CalculateBlobSize() int64 {
	size := int64(0)
	for _, f := range pd.Files {
		size += f.Blob.Size
	}
	return size
}

// GetPackageDescriptor gets the package description for a version
func GetPackageDescriptor(ctx context.Context, pv *PackageVersion) (*PackageDescriptor, error) {
	p, err := GetPackageByID(ctx, pv.PackageID)
	if err != nil {
		return nil, err
	}
	o, err := user_model.GetUserByIDCtx(ctx, p.OwnerID)
	if err != nil {
		return nil, err
	}
	repository, err := repo_model.GetRepositoryByIDCtx(ctx, p.RepoID)
	if err != nil && !repo_model.IsErrRepoNotExist(err) {
		return nil, err
	}
	creator, err := user_model.GetUserByIDCtx(ctx, pv.CreatorID)
	if err != nil {
		return nil, err
	}
	var semVer *version.Version
	if p.SemverCompatible {
		semVer, err = version.NewVersion(pv.Version)
		if err != nil {
			return nil, err
		}
	}
	pps, err := GetProperties(ctx, PropertyTypePackage, p.ID)
	if err != nil {
		return nil, err
	}
	pvps, err := GetProperties(ctx, PropertyTypeVersion, pv.ID)
	if err != nil {
		return nil, err
	}
	pfs, err := GetFilesByVersionID(ctx, pv.ID)
	if err != nil {
		return nil, err
	}

	pfds := make([]*PackageFileDescriptor, 0, len(pfs))
	for _, pf := range pfs {
		pfd, err := GetPackageFileDescriptor(ctx, pf)
		if err != nil {
			return nil, err
		}
		pfds = append(pfds, pfd)
	}

	var metadata interface{}
	switch p.Type {
	case TypeComposer:
		metadata = &composer.Metadata{}
	case TypeConan:
		metadata = &conan.Metadata{}
	case TypeContainer:
		metadata = &container.Metadata{}
	case TypeGeneric:
		// generic packages have no metadata
	case TypeHelm:
		metadata = &helm.Metadata{}
	case TypeNuGet:
		metadata = &nuget.Metadata{}
	case TypeNpm:
		metadata = &npm.Metadata{}
	case TypeMaven:
		metadata = &maven.Metadata{}
	case TypePyPI:
		metadata = &pypi.Metadata{}
	case TypeRubyGems:
		metadata = &rubygems.Metadata{}
	default:
		panic(fmt.Sprintf("unknown package type: %s", string(p.Type)))
	}
	if metadata != nil {
		if err := json.Unmarshal([]byte(pv.MetadataJSON), &metadata); err != nil {
			return nil, err
		}
	}

	return &PackageDescriptor{
		Package:           p,
		Owner:             o,
		Repository:        repository,
		Version:           pv,
		SemVer:            semVer,
		Creator:           creator,
		PackageProperties: PackagePropertyList(pps),
		VersionProperties: PackagePropertyList(pvps),
		Metadata:          metadata,
		Files:             pfds,
	}, nil
}

// GetPackageFileDescriptor gets a package file descriptor for a package file
func GetPackageFileDescriptor(ctx context.Context, pf *PackageFile) (*PackageFileDescriptor, error) {
	pb, err := GetBlobByID(ctx, pf.BlobID)
	if err != nil {
		return nil, err
	}
	pfps, err := GetProperties(ctx, PropertyTypeFile, pf.ID)
	if err != nil {
		return nil, err
	}
	return &PackageFileDescriptor{
		pf,
		pb,
		PackagePropertyList(pfps),
	}, nil
}

// GetPackageDescriptors gets the package descriptions for the versions
func GetPackageDescriptors(ctx context.Context, pvs []*PackageVersion) ([]*PackageDescriptor, error) {
	pds := make([]*PackageDescriptor, 0, len(pvs))
	for _, pv := range pvs {
		pd, err := GetPackageDescriptor(ctx, pv)
		if err != nil {
			return nil, err
		}
		pds = append(pds, pd)
	}
	return pds, nil
}