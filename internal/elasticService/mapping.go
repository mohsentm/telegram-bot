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

var AudioMapping = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"file_id":{
				"type":"keyword"
			},
			"title":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
			"caption":{
				"type":"text",
				"store": true,
				"fielddata": true
			}
		}
	}
}
`

func GetMapping(indexName string) string {
	switch indexName {
	case "twitter":
		return Twitter
	case "telegram":
		return AudioMapping
	default:
		return ""
	}
}
