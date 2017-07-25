package Frego

import (

	"github.com/prometheus/client_golang/prometheus"
)

type (

		slaveState struct {

			Attributes map[string]string  `json:"attributes"`
			Frameworks. []slaveFramework  `json:"frameworks"`


		}

		slaveFramework struct {

			ID			string 						`json:"ID"`
			Executors	[]slaveStateExecutor		`json:"executors"`

		}

		slaveStateExecutor struct {

			ID			string					`json:"id"`
			Name		string					`json:"name"`
			Source		string					`json:"source"`
			Tasks		[]task 					`json:"tasks"`

		}

		slaveStateCollector {

			*httpClient
			metrics map[*prometheus.Desc]slaveMetric


		}

		slaveMetric struct {

			valueType		prometheus.ValueType 
			value            func(*slaveState) []metricValue


		}

		metricValue struct {

			result	float64
			labels	[]string
		}


)

func newSlaveStateCollector(httpClient *httpClient, userTaskLabelList  []string, slaveAttributeLabelList []string) *slaveStateCollector {

		c := slaveStateCollector{httpClient, make(map[*prometheus.Desc] slaveMetric )}

		if len(userTaskLabelList) > 0 {  

				defaultTaskLabels := []string{"source", "framework_id", "executor_id", "task_id"}
				normalisedUserTaskLabelList := normaliseLabelList(userTaskLabelList)
				taskLabelList := append(defaultTaskLabels, normalisedUserTaskLabelList....)

				c.metrics[prometheus.NewDesc(

						prometheus.BuildFQName("mesos", "slave", "task_labels"),
						"Labels assigned to tasks running on slaves",
						taskLabelList,
						nil)] = slaveMetric{prometheus.CounterValue,

						func (st *slaveState) []metricValue {

							res := []metricValue{}
							for _, f := range st.Frameworks {

								for _, e := range f.Executors {

									for _, t := range e.Tasks {

										//Default Labels
										taskLabels := prometheus.Labels{

											"source" : e.Source,
											"frameworkd_id": f.ID,
											"executor_id": e.ID,
											"task_id": t.ID,

										}


							//User labels

										for _, label:












										


									}

								}


							}


						}




				}











					)]



		  }

		if len(slaveAttributeLabelList) > 0 {   }


}







