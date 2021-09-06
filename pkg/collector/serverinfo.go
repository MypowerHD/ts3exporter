package collector

import (
	"log"

	"github.com/hikhvar/ts3exporter/pkg/serverquery"
	"github.com/prometheus/client_golang/prometheus"
)

const serverInfoSubsystem = "serverinfo"

var serverInfoLabels = []string{virtualServerLabel}

type ServerInfo struct {
	executor        serverquery.Executor
	internalMetrics *ExporterMetrics

	ClientsOnline             *prometheus.Desc
	QueryClientsOnline        *prometheus.Desc
	Online                    *prometheus.Desc
	MaxClients                *prometheus.Desc
	Uptime                    *prometheus.Desc
	ChannelsOnline            *prometheus.Desc
	MaxDownloadTotalBandwidth *prometheus.Desc
	MaxUploadTotalBandwidth   *prometheus.Desc
	ClientsConnections        *prometheus.Desc
	QueryClientsConnections   *prometheus.Desc

	FileTransferBytesSentTotal     *prometheus.Desc
	FileTransferBytesReceivedTotal *prometheus.Desc

	ControlBytesSendTotal     *prometheus.Desc
	ControlBytesReceivedTotal *prometheus.Desc

	SpeechBytesSendTotal     *prometheus.Desc
	SpeechBytesReceivedTotal *prometheus.Desc

	KeepAliveBytesSendTotal     *prometheus.Desc
	KeepAliveBytesReceivedTotal *prometheus.Desc

	BytesSendTotal     *prometheus.Desc
	BytesReceivedTotal *prometheus.Desc

	Ping          *prometheus.Desc
	ReservedSlots *prometheus.Desc

	BandwidthReceivedLastMinuteTotal *prometheus.Desc
	BandwidthReceivedLastSecondTotal *prometheus.Desc
	BandwidthSentLastMinuteTotal     *prometheus.Desc
	BandwidthSentLastSecondTotal     *prometheus.Desc

	PacketLossSpeech    *prometheus.Desc
	PacketLossKeepAlive *prometheus.Desc
	PacketLossControl   *prometheus.Desc
	PacketLossTotal     *prometheus.Desc

	PacketsSentSpeech        *prometheus.Desc
	PacketsReceivedSpeech    *prometheus.Desc
	PacketsSentKeepalive     *prometheus.Desc
	PacketsReceivedKeepalive *prometheus.Desc
	PacketsSentControl       *prometheus.Desc
	PacketsReceivedControl   *prometheus.Desc
	PacketsSentTotal         *prometheus.Desc
	PacketsReceivedTotal     *prometheus.Desc
	BytesTotalUploaded       *prometheus.Desc
	BytesTotalDownloaded     *prometheus.Desc
}

func NewServerInfo(executor serverquery.Executor, internalMetrics *ExporterMetrics) *ServerInfo {
	return &ServerInfo{
		executor:                       executor,
		internalMetrics:                internalMetrics,
		ClientsOnline:                  prometheus.NewDesc(fqdn(serverInfoSubsystem, "clients_online"), "number of currently online clients", serverInfoLabels, nil),
		QueryClientsOnline:             prometheus.NewDesc(fqdn(serverInfoSubsystem, "query_clients_online"), "number of currently online query clients", serverInfoLabels, nil),
		Online:                         prometheus.NewDesc(fqdn(serverInfoSubsystem, "online"), "is the virtual server online", serverInfoLabels, nil),
		MaxClients:                     prometheus.NewDesc(fqdn(serverInfoSubsystem, "max_clients"), "maximal number of allowed clients", serverInfoLabels, nil),
		Uptime:                         prometheus.NewDesc(fqdn(serverInfoSubsystem, "uptime"), "uptime of the virtual server", serverInfoLabels, nil),
		ChannelsOnline:                 prometheus.NewDesc(fqdn(serverInfoSubsystem, "channels_online"), "number of online channels", serverInfoLabels, nil),
		MaxDownloadTotalBandwidth:      prometheus.NewDesc(fqdn(serverInfoSubsystem, "download_bandwidth_bytes_max"), "maximal bandwidth available for downloads", serverInfoLabels, nil),
		MaxUploadTotalBandwidth:        prometheus.NewDesc(fqdn(serverInfoSubsystem, "upload_bandwidth_bytes_max"), "maximal bandwidth available for uploads", serverInfoLabels, nil),
		ClientsConnections:             prometheus.NewDesc(fqdn(serverInfoSubsystem, "client_connections"), "currently established client connections", serverInfoLabels, nil),
		QueryClientsConnections:        prometheus.NewDesc(fqdn(serverInfoSubsystem, "query_client_connections"), "currently established query client connections", serverInfoLabels, nil),
		FileTransferBytesSentTotal:     prometheus.NewDesc(fqdn(serverInfoSubsystem, "file_transfer_bytes_sent_total"), "total sent bytes for file transfers", serverInfoLabels, nil),
		FileTransferBytesReceivedTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "file_transfer_bytes_received_total"), "total received bytes for file transfers", serverInfoLabels, nil),
		ControlBytesSendTotal:          prometheus.NewDesc(fqdn(serverInfoSubsystem, "control_bytes_sent_total"), "total sent bytes for control traffic", serverInfoLabels, nil),
		ControlBytesReceivedTotal:      prometheus.NewDesc(fqdn(serverInfoSubsystem, "control_bytes_received_total"), "total received bytes for control traffic", serverInfoLabels, nil),
		SpeechBytesSendTotal:           prometheus.NewDesc(fqdn(serverInfoSubsystem, "speech_bytes_sent_total"), "total sent bytes for speech traffic", serverInfoLabels, nil),
		SpeechBytesReceivedTotal:       prometheus.NewDesc(fqdn(serverInfoSubsystem, "speech_bytes_received_total"), "total received bytes for speech traffic", serverInfoLabels, nil),
		KeepAliveBytesSendTotal:        prometheus.NewDesc(fqdn(serverInfoSubsystem, "keepalive_bytes_sent_total"), "total send bytes for keepalive traffic", serverInfoLabels, nil),
		KeepAliveBytesReceivedTotal:    prometheus.NewDesc(fqdn(serverInfoSubsystem, "keepalive_bytes_received_total"), "total received bytes for keepalive traffic", serverInfoLabels, nil),
		BytesSendTotal:                 prometheus.NewDesc(fqdn(serverInfoSubsystem, "bytes_send_total"), "total send bytes", serverInfoLabels, nil),
		BytesReceivedTotal:             prometheus.NewDesc(fqdn(serverInfoSubsystem, "bytes_received_total"), "total received bytes", serverInfoLabels, nil),

		Ping:          prometheus.NewDesc(fqdn(serverInfoSubsystem, "ping_total"), "Total ping", serverInfoLabels, nil),
		ReservedSlots: prometheus.NewDesc(fqdn(serverInfoSubsystem, "reserved_slots"), "How many Slots are Reserved", serverInfoLabels, nil),

		BandwidthReceivedLastMinuteTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "bandwidth_received_last_minute_total"), "Last minutes Total Recived", serverInfoLabels, nil),
		BandwidthReceivedLastSecondTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "bandwidth_received_last_second_total"), "Last Second Total Recived", serverInfoLabels, nil),
		BandwidthSentLastMinuteTotal:     prometheus.NewDesc(fqdn(serverInfoSubsystem, "bandwidth_sent_last_minute_total"), "Last Minute Sent Total", serverInfoLabels, nil),
		BandwidthSentLastSecondTotal:     prometheus.NewDesc(fqdn(serverInfoSubsystem, "bandwidth_sent_last_second_total"), "Last Second Sent Total", serverInfoLabels, nil),

		PacketLossSpeech:    prometheus.NewDesc(fqdn(serverInfoSubsystem, "total_packetloss_speech"), "Total Packetloss Speech", serverInfoLabels, nil),
		PacketLossKeepAlive: prometheus.NewDesc(fqdn(serverInfoSubsystem, "total_packetloss_keepalive"), "Total Packetloss Keepalive", serverInfoLabels, nil),
		//total_packetloss_control
		PacketLossControl: prometheus.NewDesc(fqdn(serverInfoSubsystem, "total_packetloss_control"), "", serverInfoLabels, nil),
		//total_packetloss_total
		PacketLossTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packetloss_total"), "Total Packet Lost", serverInfoLabels, nil),
		//packets_sent_speech
		PacketsSentSpeech: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_sent_speech"), "Packets Sent Speech", serverInfoLabels, nil),
		//packets_received_speech
		PacketsReceivedSpeech: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_received_speech"), "Packets received Speech", serverInfoLabels, nil),
		//packets_sent_keepalive
		PacketsSentKeepalive: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_sent_keepalive"), "Packets sent Keepalive", serverInfoLabels, nil),
		//packets_received_keepalive
		PacketsReceivedKeepalive: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_received_keepalive"), "Packets received Keepalive", serverInfoLabels, nil),
		//packets_sent_control
		PacketsSentControl: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_sent_control"), "Packets sent Control", serverInfoLabels, nil),
		//packets_received_control
		PacketsReceivedControl: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_received_control"), "Packets received Control", serverInfoLabels, nil),
		//packets_sent_total
		PacketsSentTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_sent_total"), "Packets sent Total", serverInfoLabels, nil),
		//packets_received_total
		PacketsReceivedTotal: prometheus.NewDesc(fqdn(serverInfoSubsystem, "packets_received_total"), "Packets received Total", serverInfoLabels, nil),

		//total_bytes_uploaded
		BytesTotalUploaded: prometheus.NewDesc(fqdn(serverInfoSubsystem, "total_bytes_uploaded"), "Bytes Total Uploaded", serverInfoLabels, nil),
		//total_bytes_downloaded
		BytesTotalDownloaded: prometheus.NewDesc(fqdn(serverInfoSubsystem, "total_bytes_downloaded"), "Bytes Total Uploaded", serverInfoLabels, nil),
	}

}

func (s *ServerInfo) Describe(c chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(s, c)
}

func (s *ServerInfo) Collect(c chan<- prometheus.Metric) {
	vServerView := serverquery.NewVirtualServer(s.executor)
	if err := vServerView.Refresh(); err != nil {
		s.internalMetrics.RefreshError(serverInfoSubsystem)
		log.Printf("failed to refresh server info view: %v", err)
	}
	for _, vs := range vServerView.All() {
		c <- prometheus.MustNewConstMetric(s.ClientsOnline, prometheus.GaugeValue, float64(vs.ClientsOnline), vs.Name)
		c <- prometheus.MustNewConstMetric(s.QueryClientsOnline, prometheus.GaugeValue, float64(vs.QueryClientsOnline), vs.Name)
		c <- prometheus.MustNewConstMetric(s.Online, prometheus.GaugeValue, online(vs.Status), vs.Name)
		c <- prometheus.MustNewConstMetric(s.MaxClients, prometheus.GaugeValue, float64(vs.MaxClients), vs.Name)
		c <- prometheus.MustNewConstMetric(s.Uptime, prometheus.CounterValue, float64(vs.Uptime), vs.Name)
		c <- prometheus.MustNewConstMetric(s.ChannelsOnline, prometheus.GaugeValue, float64(vs.ChannelsOnline), vs.Name)
		c <- prometheus.MustNewConstMetric(s.MaxDownloadTotalBandwidth, prometheus.GaugeValue, float64(vs.MaxDownloadTotalBandwidth), vs.Name)
		c <- prometheus.MustNewConstMetric(s.MaxUploadTotalBandwidth, prometheus.GaugeValue, float64(vs.MaxUploadTotalBandwidth), vs.Name)
		c <- prometheus.MustNewConstMetric(s.ClientsConnections, prometheus.GaugeValue, float64(vs.ClientsConnections), vs.Name)
		c <- prometheus.MustNewConstMetric(s.QueryClientsConnections, prometheus.GaugeValue, float64(vs.QueryClientsConnections), vs.Name)
		c <- prometheus.MustNewConstMetric(s.FileTransferBytesSentTotal, prometheus.CounterValue, float64(vs.FileTransferBytesSentTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.FileTransferBytesReceivedTotal, prometheus.CounterValue, float64(vs.FileTransferBytesReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.ControlBytesSendTotal, prometheus.CounterValue, float64(vs.ControlBytesSendTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.ControlBytesReceivedTotal, prometheus.CounterValue, float64(vs.ControlBytesReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.SpeechBytesSendTotal, prometheus.CounterValue, float64(vs.SpeechBytesSendTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.SpeechBytesReceivedTotal, prometheus.CounterValue, float64(vs.SpeechBytesReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.KeepAliveBytesSendTotal, prometheus.CounterValue, float64(vs.KeepAliveBytesSendTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.KeepAliveBytesReceivedTotal, prometheus.CounterValue, float64(vs.KeepAliveBytesReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BytesSendTotal, prometheus.CounterValue, float64(vs.BytesSendTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BytesReceivedTotal, prometheus.CounterValue, float64(vs.BytesReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.Ping, prometheus.CounterValue, vs.Ping, vs.Name)
		c <- prometheus.MustNewConstMetric(s.ReservedSlots, prometheus.CounterValue, float64(vs.ReservedSlots), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BandwidthReceivedLastMinuteTotal, prometheus.CounterValue, float64(vs.BandwidthReceivedLastMinuteTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BandwidthReceivedLastSecondTotal, prometheus.CounterValue, float64(vs.BandwidthReceivedLastSecondTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BandwidthSentLastMinuteTotal, prometheus.CounterValue, float64(vs.BandwidthSentLastMinuteTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BandwidthSentLastSecondTotal, prometheus.CounterValue, float64(vs.BandwidthSentLastSecondTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketLossSpeech, prometheus.CounterValue, vs.PacketLossSpeech, vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketLossKeepAlive, prometheus.CounterValue, vs.PacketLossKeepAlive, vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketLossControl, prometheus.CounterValue, vs.PacketLossControl, vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketLossTotal, prometheus.CounterValue, vs.PacketLossTotal, vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsSentSpeech, prometheus.CounterValue, float64(vs.PacketsSentSpeech), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsReceivedSpeech, prometheus.CounterValue, float64(vs.PacketsReceivedSpeech), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsSentKeepalive, prometheus.CounterValue, float64(vs.PacketsSentKeepalive), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsReceivedKeepalive, prometheus.CounterValue, float64(vs.PacketsReceivedKeepalive), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsSentControl, prometheus.CounterValue, float64(vs.PacketsSentControl), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsReceivedControl, prometheus.CounterValue, float64(vs.PacketsReceivedControl), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsSentTotal, prometheus.CounterValue, float64(vs.PacketsSentTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.PacketsReceivedTotal, prometheus.CounterValue, float64(vs.PacketsReceivedTotal), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BytesTotalUploaded, prometheus.CounterValue, float64(vs.BytesTotalUploaded), vs.Name)
		c <- prometheus.MustNewConstMetric(s.BytesTotalDownloaded, prometheus.CounterValue, float64(vs.BytesTotalDownloaded), vs.Name)
	}
}

func online(status string) float64 {
	if status == "online" {
		return 1.0
	}
	return 0.0
}
