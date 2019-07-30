// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var PeerInfoFieldPathsNested = []string{
	"grpc_port",
	"roles",
	"tags",
	"tls",
}

var PeerInfoFieldPathsTopLevel = []string{
	"grpc_port",
	"roles",
	"tags",
	"tls",
}
var ClusterFieldPathsNested = []string{
	"addresses",
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"endpoints",
	"ids",
	"ids.cluster_id",
	"location",
	"location.accuracy",
	"location.altitude",
	"location.latitude",
	"location.longitude",
	"location.source",
	"location_description",
	"name",
	"roles",
	"secret",
	"updated_at",
}

var ClusterFieldPathsTopLevel = []string{
	"addresses",
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"endpoints",
	"ids",
	"location",
	"location_description",
	"name",
	"roles",
	"secret",
	"updated_at",
}
var ClustersFieldPathsNested = []string{
	"clusters",
}

var ClustersFieldPathsTopLevel = []string{
	"clusters",
}
var GetClusterRequestFieldPathsNested = []string{
	"cluster_ids",
	"cluster_ids.cluster_id",
	"field_mask",
}

var GetClusterRequestFieldPathsTopLevel = []string{
	"cluster_ids",
	"field_mask",
}
var GetClusterIdentifiersForAddressRequestFieldPathsNested = []string{
	"address",
}

var GetClusterIdentifiersForAddressRequestFieldPathsTopLevel = []string{
	"address",
}
var ListClustersRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"field_mask",
	"limit",
	"order",
	"page",
}

var ListClustersRequestFieldPathsTopLevel = []string{
	"collaborator",
	"field_mask",
	"limit",
	"order",
	"page",
}
var CreateClusterRequestFieldPathsNested = []string{
	"cluster",
	"cluster.addresses",
	"cluster.attributes",
	"cluster.contact_info",
	"cluster.created_at",
	"cluster.description",
	"cluster.endpoints",
	"cluster.ids",
	"cluster.ids.cluster_id",
	"cluster.location",
	"cluster.location.accuracy",
	"cluster.location.altitude",
	"cluster.location.latitude",
	"cluster.location.longitude",
	"cluster.location.source",
	"cluster.location_description",
	"cluster.name",
	"cluster.roles",
	"cluster.secret",
	"cluster.updated_at",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
}

var CreateClusterRequestFieldPathsTopLevel = []string{
	"cluster",
	"collaborator",
}
var UpdateClusterRequestFieldPathsNested = []string{
	"cluster",
	"cluster.addresses",
	"cluster.attributes",
	"cluster.contact_info",
	"cluster.created_at",
	"cluster.description",
	"cluster.endpoints",
	"cluster.ids",
	"cluster.ids.cluster_id",
	"cluster.location",
	"cluster.location.accuracy",
	"cluster.location.altitude",
	"cluster.location.latitude",
	"cluster.location.longitude",
	"cluster.location.source",
	"cluster.location_description",
	"cluster.name",
	"cluster.roles",
	"cluster.secret",
	"cluster.updated_at",
	"field_mask",
}

var UpdateClusterRequestFieldPathsTopLevel = []string{
	"cluster",
	"field_mask",
}
var ListClusterCollaboratorsRequestFieldPathsNested = []string{
	"cluster_ids",
	"cluster_ids.cluster_id",
	"limit",
	"page",
}

var ListClusterCollaboratorsRequestFieldPathsTopLevel = []string{
	"cluster_ids",
	"limit",
	"page",
}
var GetClusterCollaboratorRequestFieldPathsNested = []string{
	"cluster_ids",
	"cluster_ids.cluster_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
}

var GetClusterCollaboratorRequestFieldPathsTopLevel = []string{
	"cluster_ids",
	"collaborator",
}
var SetClusterCollaboratorRequestFieldPathsNested = []string{
	"cluster_ids",
	"cluster_ids.cluster_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
}

var SetClusterCollaboratorRequestFieldPathsTopLevel = []string{
	"cluster_ids",
	"collaborator",
}
var Cluster_EndpointFieldPathsNested = []string{
	"endpoint",
	"endpoint.backend_interfaces_http",
	"endpoint.backend_interfaces_http.host",
	"endpoint.backend_interfaces_http.path",
	"endpoint.backend_interfaces_http.port",
	"endpoint.backend_interfaces_http.tls",
	"endpoint.basic_station_http",
	"endpoint.basic_station_http.host",
	"endpoint.basic_station_http.path",
	"endpoint.basic_station_http.port",
	"endpoint.basic_station_http.tls",
	"endpoint.grpc",
	"endpoint.grpc.host",
	"endpoint.grpc.port",
	"endpoint.grpc.tls",
	"endpoint.grpc_http",
	"endpoint.grpc_http.host",
	"endpoint.grpc_http.path",
	"endpoint.grpc_http.port",
	"endpoint.grpc_http.tls",
	"endpoint.http",
	"endpoint.http.host",
	"endpoint.http.path",
	"endpoint.http.port",
	"endpoint.http.tls",
	"endpoint.mqtt",
	"endpoint.mqtt.host",
	"endpoint.mqtt.port",
	"endpoint.mqtt.tls",
	"endpoint.packet_forwarder_udp",
	"endpoint.packet_forwarder_udp.dtls",
	"endpoint.packet_forwarder_udp.host",
	"endpoint.packet_forwarder_udp.port",
	"roles",
}

var Cluster_EndpointFieldPathsTopLevel = []string{
	"endpoint",
	"roles",
}
var Cluster_Endpoint_GRPCFieldPathsNested = []string{
	"host",
	"port",
	"tls",
}

var Cluster_Endpoint_GRPCFieldPathsTopLevel = []string{
	"host",
	"port",
	"tls",
}
var Cluster_Endpoint_HTTPFieldPathsNested = []string{
	"host",
	"path",
	"port",
	"tls",
}

var Cluster_Endpoint_HTTPFieldPathsTopLevel = []string{
	"host",
	"path",
	"port",
	"tls",
}
var Cluster_Endpoint_MQTTFieldPathsNested = []string{
	"host",
	"port",
	"tls",
}

var Cluster_Endpoint_MQTTFieldPathsTopLevel = []string{
	"host",
	"port",
	"tls",
}
var Cluster_Endpoint_UDPFieldPathsNested = []string{
	"dtls",
	"host",
	"port",
}

var Cluster_Endpoint_UDPFieldPathsTopLevel = []string{
	"dtls",
	"host",
	"port",
}
