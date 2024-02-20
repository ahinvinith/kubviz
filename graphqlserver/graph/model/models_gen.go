// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cluster struct {
	Name string `json:"name"`
}

type ClusterAPIsCount struct {
	ClusterName string `json:"clusterName"`
	Count       int    `json:"count"`
}

type ClusterNamespaceOutdatedCount struct {
	ClusterName   string `json:"clusterName"`
	Namespace     string `json:"namespace"`
	OutdatedCount int    `json:"outdatedCount"`
}

type ClusterNamespaceResourceCount struct {
	ClusterName   string `json:"clusterName"`
	Namespace     string `json:"namespace"`
	ResourceCount int    `json:"resourceCount"`
}

type DeletedAPI struct {
	ClusterName *string `json:"ClusterName,omitempty"`
	ObjectName  *string `json:"ObjectName,omitempty"`
	Group       *string `json:"Group,omitempty"`
	Kind        *string `json:"Kind,omitempty"`
	Version     *string `json:"Version,omitempty"`
	Name        *string `json:"Name,omitempty"`
	Deleted     *bool   `json:"Deleted,omitempty"`
	Scope       *string `json:"Scope,omitempty"`
	EventTime   *string `json:"EventTime,omitempty"`
	ExpiryDate  *string `json:"ExpiryDate,omitempty"`
}

type DeprecatedAPI struct {
	ClusterName *string `json:"ClusterName,omitempty"`
	ObjectName  *string `json:"ObjectName,omitempty"`
	Description *string `json:"Description,omitempty"`
	Kind        *string `json:"Kind,omitempty"`
	Deprecated  *bool   `json:"Deprecated,omitempty"`
	Scope       *string `json:"Scope,omitempty"`
	EventTime   *string `json:"EventTime,omitempty"`
	ExpiryDate  *string `json:"ExpiryDate,omitempty"`
}

type Event struct {
	ClusterName *string `json:"ClusterName,omitempty"`
	ID          *string `json:"Id,omitempty"`
	EventTime   *string `json:"EventTime,omitempty"`
	OpType      *string `json:"OpType,omitempty"`
	Name        *string `json:"Name,omitempty"`
	Namespace   *string `json:"Namespace,omitempty"`
	Kind        *string `json:"Kind,omitempty"`
	Message     *string `json:"Message,omitempty"`
	Reason      *string `json:"Reason,omitempty"`
	Host        *string `json:"Host,omitempty"`
	Event       *string `json:"Event,omitempty"`
	ImageName   *string `json:"ImageName,omitempty"`
	FirstTime   *string `json:"FirstTime,omitempty"`
	LastTime    *string `json:"LastTime,omitempty"`
	ExpiryDate  *string `json:"ExpiryDate,omitempty"`
}

type GetAllResource struct {
	ClusterName *string `json:"ClusterName,omitempty"`
	Namespace   *string `json:"Namespace,omitempty"`
	Kind        *string `json:"Kind,omitempty"`
	Resource    *string `json:"Resource,omitempty"`
	Age         *string `json:"Age,omitempty"`
	EventTime   *string `json:"EventTime,omitempty"`
	ExpiryDate  *string `json:"ExpiryDate,omitempty"`
}

type KubeScore struct {
	ID          string `json:"id"`
	ClusterName string `json:"clusterName"`
	ObjectName  string `json:"objectName"`
	Kind        string `json:"kind"`
	APIVersion  string `json:"apiVersion"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	TargetType  string `json:"targetType"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	FileName    string `json:"fileName"`
	FileRow     int    `json:"fileRow"`
	EventTime   string `json:"eventTime"`
}

type Kubescore struct {
	ID          string  `json:"id"`
	ClusterName *string `json:"clusterName,omitempty"`
	ObjectName  *string `json:"objectName,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	APIVersion  *string `json:"apiVersion,omitempty"`
	Name        *string `json:"name,omitempty"`
	Namespace   *string `json:"namespace,omitempty"`
	TargetType  *string `json:"targetType,omitempty"`
	Description *string `json:"description,omitempty"`
	Path        *string `json:"path,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	FileName    *string `json:"fileName,omitempty"`
	FileRow     *int    `json:"fileRow,omitempty"`
	EventTime   *string `json:"eventTime,omitempty"`
	ExpiryDate  *string `json:"expiryDate,omitempty"`
}

type Namespace struct {
	Name string `json:"name"`
}

type NamespaceData struct {
	Namespace      string           `json:"namespace"`
	OutdatedImages []*OutdatedImage `json:"outdatedImages"`
	KubeScores     []*KubeScore     `json:"kubeScores"`
	Resources      []*Resource      `json:"resources"`
}

type OutdatedImage struct {
	ClusterName    string `json:"clusterName"`
	Namespace      string `json:"namespace"`
	Pod            string `json:"pod"`
	CurrentImage   string `json:"currentImage"`
	CurrentTag     string `json:"currentTag"`
	LatestVersion  string `json:"latestVersion"`
	VersionsBehind int    `json:"versionsBehind"`
	EventTime      string `json:"eventTime"`
}

type Query struct {
}

type Rakkess struct {
	ClusterName *string `json:"ClusterName,omitempty"`
	Name        *string `json:"Name,omitempty"`
	Create      *string `json:"Create,omitempty"`
	Delete      *string `json:"Delete,omitempty"`
	List        *string `json:"List,omitempty"`
	Update      *string `json:"Update,omitempty"`
	EventTime   *string `json:"EventTime,omitempty"`
	ExpiryDate  *string `json:"ExpiryDate,omitempty"`
}

type Resource struct {
	ClusterName string `json:"clusterName"`
	Namespace   string `json:"namespace"`
	Kind        string `json:"kind"`
	Resource    string `json:"resource"`
	Age         string `json:"age"`
	EventTime   string `json:"eventTime"`
}

type TrivyImage struct {
	ID                  string  `json:"id"`
	ClusterName         *string `json:"clusterName,omitempty"`
	ArtifactName        *string `json:"artifactName,omitempty"`
	VulID               *string `json:"vulId,omitempty"`
	VulPkgID            *string `json:"vulPkgId,omitempty"`
	VulPkgName          *string `json:"vulPkgName,omitempty"`
	VulInstalledVersion *string `json:"vulInstalledVersion,omitempty"`
	VulFixedVersion     *string `json:"vulFixedVersion,omitempty"`
	VulTitle            *string `json:"vulTitle,omitempty"`
	VulSeverity         *string `json:"vulSeverity,omitempty"`
	VulPublishedDate    *string `json:"vulPublishedDate,omitempty"`
	VulLastModifiedDate *string `json:"vulLastModifiedDate,omitempty"`
	ExpiryDate          *string `json:"expiryDate,omitempty"`
}

type TrivyMisconfig struct {
	ID                  string  `json:"id"`
	ClusterName         *string `json:"clusterName,omitempty"`
	Namespace           *string `json:"namespace,omitempty"`
	Kind                *string `json:"kind,omitempty"`
	Name                *string `json:"name,omitempty"`
	MisconfigID         *string `json:"misconfigId,omitempty"`
	MisconfigAvdid      *string `json:"misconfigAvdid,omitempty"`
	MisconfigType       *string `json:"misconfigType,omitempty"`
	MisconfigTitle      *string `json:"misconfigTitle,omitempty"`
	MisconfigDesc       *string `json:"misconfigDesc,omitempty"`
	MisconfigMsg        *string `json:"misconfigMsg,omitempty"`
	MisconfigQuery      *string `json:"misconfigQuery,omitempty"`
	MisconfigResolution *string `json:"misconfigResolution,omitempty"`
	MisconfigSeverity   *string `json:"misconfigSeverity,omitempty"`
	MisconfigStatus     *string `json:"misconfigStatus,omitempty"`
	EventTime           *string `json:"eventTime,omitempty"`
	ExpiryDate          *string `json:"expiryDate,omitempty"`
}

type TrivySbom struct {
	ID           string  `json:"id"`
	ClusterName  *string `json:"clusterName,omitempty"`
	ImageName    *string `json:"imageName,omitempty"`
	PackageName  *string `json:"packageName,omitempty"`
	PackageURL   *string `json:"packageUrl,omitempty"`
	BomRef       *string `json:"bomRef,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Version      *int    `json:"version,omitempty"`
	BomFormat    *string `json:"bomFormat,omitempty"`
	ExpiryDate   *string `json:"expiryDate,omitempty"`
}

type TrivyVul struct {
	ID                  string  `json:"id"`
	ClusterName         *string `json:"clusterName,omitempty"`
	Namespace           *string `json:"namespace,omitempty"`
	Kind                *string `json:"kind,omitempty"`
	Name                *string `json:"name,omitempty"`
	VulID               *string `json:"vulId,omitempty"`
	VulVendorIds        *string `json:"vulVendorIds,omitempty"`
	VulPkgID            *string `json:"vulPkgId,omitempty"`
	VulPkgName          *string `json:"vulPkgName,omitempty"`
	VulPkgPath          *string `json:"vulPkgPath,omitempty"`
	VulInstalledVersion *string `json:"vulInstalledVersion,omitempty"`
	VulFixedVersion     *string `json:"vulFixedVersion,omitempty"`
	VulTitle            *string `json:"vulTitle,omitempty"`
	VulSeverity         *string `json:"vulSeverity,omitempty"`
	VulPublishedDate    *string `json:"vulPublishedDate,omitempty"`
	VulLastModifiedDate *string `json:"vulLastModifiedDate,omitempty"`
	ExpiryDate          *string `json:"expiryDate,omitempty"`
}
