package elasticservice

var Twitter = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"user":{
				"type":"keyword"
			},
			"message":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
			"retweets":{
				"type":"text"
			},
			"tags":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			},
			"suggest_field":{
				"type":"completion"
			}
		}
	}
}
`

var Audio = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"fileID":{
				"type":"keyword"
			},
			"title":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
		}
	}
}
`

func GetMapping() string {
	return Twitter
}
