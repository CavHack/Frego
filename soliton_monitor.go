package Frego

import (

		"github.com/prometheus/client_golang/prometheus"
)

type (

		executor struct {


				ID       			string 			`json:"executor_id"`
				Name	 			string			`json:"executor_name"`
				FrameworkID			string			`json:"framework_id"`
				Source				string			`json:"source"`
				Statistics			*statistics		`json:"statistics"`
				Tasks				[]task 			`json:"tasks"`	



		}


		statistics   struct {

			Processes					float64			`json:"processes"`
			Threads						float64			`json:"threads"`

			CpusLimit					float64			`json:"cpus_limit"`
			CpusSystemTimeSecs  		float64			`json:"cpus_system_time_secs"`
			CpusUserTimeSecs			float64			`json:"cpus_user_time_secs"`
			CpusThrottledTimeSecs		float64			`json:"cpus_throttled_time_secs"`
			CpusNrPeriods				float64			`json:"cpus_nr_periods"`
			CpusNrThrottled				float64			`json:cpus_nr_throttled`

			MemLimitBytes				float64			`json:"mem_limit_bytes"`
			MemRssBytes					float64			`json:"mem_rss_bytes"`
			MemTotalBytes				float64			`json:"mem_total_bytes"`
			MemCacheBytes				float64			`json:"mem_cache_bytes"`	
			MemSwapBytes				float64			`json:"mem_swap_bytes"`

			DiskLimitBytes				float64			`json:"disk_limit_bytes"`
			DiskUsedBytes				float64			`json:"disk_used_bytes"`

			NetRxBytes					float64			`json:"net_rx_bytes"`
			NetRxDropped				float64			`json:"net_rx_dropped"`
			NetRxErrors					float64			`json:"net_rx_errors"`
			NetRxPackets				float64			`json:"net_rx_packets"`
			NetTxBytes					float64			`json:"net_tx_bytes"`
			NetTxErrors					float64			`json:"net_tx_errors"`
			NetTxPackets				float64			`json:"net_tx_packets"`




		}



		slaveCollector struct {

				*httpClient
				metrics map[*prometheus.Desc]metric


		}

		metric struct {

			valueType prometheus.ValueType
			get 		func(*statistics) float64


		}



)

func newSlaveMonitorCollector(http *httpClient) prometheus.Collector {


			labels := []string{
				"id",
				"framework_id",

				"source"
			}



			return &slaveCollector{


					httpClient: httpClient,
					metrics: map[*prometheus.Desc]metric {

						//Processes
						prometheus.NewDesc(

								"processes",
								"Current Number of processes",
								labels, nil,

							): metric{prometheus.GaugeValue, func(s *statistics) float64 {return s.Processes}},
						prometheus.NewDesc(
								"threads",
								"Current number of threads",
								labels, nil,
							): metric{prometheus.GaugeValue, func(s *statistics) float64 { return s.Threads}},


						//CPU

						prometheus.NewDesc(

								"cpus_limit",
								"Current limit of CPUS for task",
								labels, nil,


							): metric{[prometheus.GaugeValue, func(s *statistics) float64 { return s.CpusLimit}]},
						prometheus.NewDesc(

								"cpus_user_seconds_total",
								"Total user CPU seconds",
								labels, nil

							)

					}



			}




}














