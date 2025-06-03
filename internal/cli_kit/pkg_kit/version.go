package pkg_kit

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

var buildInfo *BuildInfo

func SaveBuildInfo(bdInfo *BuildInfo) {
	buildInfo = bdInfo
}

func GetNewBuildInfo() BuildInfo {
	checkPackageJsonLoad()

	if buildInfo == nil {
		panic("buildInfo is nil, please init by method SaveBuildInfo()")
	}

	return *buildInfo
}

func FetchNowPkgName() string {
	checkPackageJsonLoad()

	if buildInfo == nil {
		panic("buildInfo is nil, please init by method SaveBuildInfo()")
	}

	return buildInfo.PkgName
}

func FetchNowVersion() string {
	checkPackageJsonLoad()

	if buildInfo == nil {
		panic("buildInfo is nil, please init by method SaveBuildInfo()")
	}

	return buildInfo.Version
}

func FetchNowBuildId() string {
	checkPackageJsonLoad()

	if buildInfo == nil {
		panic("buildInfo is nil, please init by method SaveBuildInfo()")
	}

	if buildInfo.BuildId == "" {
		panic("nowBuildId is empty, please init by method SaveBuildInfo()")
	}

	return buildInfo.BuildId
}

func FetchNowBuildIdShort() string {
	checkPackageJsonLoad()

	if buildInfo.BuildIdShort == "" {
		panic("nowBuildIdShort is empty, please init by method SaveBuildInfo()")
	}

	return buildInfo.BuildIdShort
}

func FetchNowBuildCode() string {
	if FetchNowBuildId() != InfoUnknown {
		return FetchNowBuildId()
	}

	return buildInfo.BuildSum
}

const (
	InfoUnknown      = "unknown"
	versionDevel     = "devel"
	buildIdShortSize = 8
)

type BuildInfo struct {
	PkgName     string `json:"pkgName"`
	Description string `json:"description"`

	Version      string `json:"version"`
	RawVersion   string `json:"rawVersion"`
	BuildId      string `json:"buildId"`
	BuildIdShort string
	BuildSum     string `json:"buildSum"`

	GoVersion    string `json:"goVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	Date         string `json:"date"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`

	AuthorName         string `json:"authorName"`
	CopyrightStartYear string `json:"copyrightStartYear"`
	CopyrightNowYear   string `json:"copyrightNowYear"`
}

func (b BuildInfo) String() string {
	return fmt.Sprintf(
		"%s has version %s, © %s-%s %s,  built with %s id: %s from %s on %s, run on %s",
		b.PkgName,
		b.Version,
		b.CopyrightStartYear,
		b.CopyrightNowYear,
		b.AuthorName,
		b.GoVersion,
		b.BuildId,
		b.GitCommit,
		b.Date,
		b.Platform,
	)
}

func (b BuildInfo) Copyright() string {
	return fmt.Sprintf("© %s-%s by: %s  build with %s id: %s, run on %s",
		b.CopyrightStartYear, b.CopyrightNowYear, b.AuthorName, b.GoVersion, b.BuildId, b.Platform)
}

func (b BuildInfo) RunInfoString() string {
	return fmt.Sprintf("%s, run on %s, built with %s id: %s from %s on %s",
		b.PkgName, b.Platform, b.GoVersion, b.BuildId, b.GitCommit, b.Date,
	)
}

func (b BuildInfo) PgkNameString() string {
	return b.PkgName
}

func (b BuildInfo) PgkFullInfo() string {
	return fmt.Sprintf("%s by: %s build with %s id: %s, run on %s",
		b.PkgName, b.AuthorName, b.GoVersion, b.BuildId, b.Platform)
}

func (b BuildInfo) DescriptionString() string {
	return b.Description
}

func (b BuildInfo) VersionString() string {
	return b.Version
}

func (b BuildInfo) RawVersionString() string {
	return b.RawVersion
}

func NewBuildInfo(
	pkgName, description,
	version, rawVersion,
	buildId, commit, date,
	author, copyrightStartYear string,
) BuildInfo {
	info := BuildInfo{
		PkgName:     pkgName,
		Description: description,

		Version:    version,
		RawVersion: rawVersion,
		BuildId:    buildId,
		GitCommit:  commit,
		Date:       date,
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),

		AuthorName:         author,
		CopyrightStartYear: copyrightStartYear,
		CopyrightNowYear:   strconv.Itoa(time.Now().Year()),
	}

	nowBuildIdShort := info.BuildId
	if len(nowBuildIdShort) > buildIdShortSize {
		nowBuildIdShort = info.BuildId[buildIdShortSize:]
	}

	info.BuildIdShort = nowBuildIdShort

	bi, available := debug.ReadBuildInfo()
	if !available {
		return info
	}

	info.GoVersion = bi.GoVersion
	if info.GoVersion == "" {
		info.GoVersion = InfoUnknown
	}

	if info.Version == "" || info.Version == InfoUnknown {
		info.Version = firstNonEmpty(getGitVersion(bi), versionDevel)
	}

	if date != "" {
		return info
	}

	var revision string

	var modified string

	for _, setting := range bi.Settings {
		// The `vcs.xxx` information is only available with `go build`.
		// This information is not available with `go install` or `go run`.
		switch setting.Key {
		case "vcs.time":
			info.Date = setting.Value
		case "vcs.revision":
			revision = setting.Value
		case "vcs.modified":
			modified = setting.Value
		}
	}

	if revision == "" {
		revision = InfoUnknown
	}

	if modified == "" {
		modified = "?"
	}

	if info.Date == "" {
		info.Date = fmt.Sprintf("(%s)", InfoUnknown)
	}

	if info.BuildId == "" {
		info.BuildId = fmt.Sprintf("(%s)", InfoUnknown)
	}

	info.GitCommit = fmt.Sprintf("(%s, modified: %s, mod sum: %q)", revision, modified, bi.Main.Sum)

	info.BuildSum = bi.Main.Sum

	return info
}

func getGitVersion(bi *debug.BuildInfo) string {
	if bi == nil {
		return ""
	}

	// remove this when the issue https://github.com/golang/go/issues/29228 is fixed
	if bi.Main.Version == "(devel)" || bi.Main.Version == "" {
		return ""
	}

	return bi.Main.Version
}

//nolint:golint,unused
func getCommit(bi *debug.BuildInfo) string {
	return getKey(bi, "vcs.revision")
}

//nolint:golint,unused
func getDirty(bi *debug.BuildInfo) string {
	modified := getKey(bi, "vcs.modified")
	if modified == "true" {
		return "dirty"
	}

	if modified == "false" {
		return "clean"
	}

	return ""
}

//nolint:golint,unused
func getBuildDate(bi *debug.BuildInfo) string {
	buildTime := getKey(bi, "vcs.time")

	t, err := time.Parse("2006-01-02T15:04:05Z", buildTime)
	if err != nil {
		return ""
	}

	return t.Format("2006-01-02T15:04:05")
}

//nolint:golint,unused
func getKey(bi *debug.BuildInfo, key string) string {
	if bi == nil {
		return ""
	}

	for _, iter := range bi.Settings {
		if iter.Key == key {
			return iter.Value
		}
	}

	return ""
}

//nolint:golint,unused
func firstNonEmpty(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}

	return ""
}
