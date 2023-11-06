package config

const SearchProjectsByUserQuery = `{
	"query": {
	  "bool": {
		"must": [
		  {
			"nested": {
			  "path": "users",
			  "query": {
				"term": {
				  "users.userId": "%s"
				}
			  }
			}
		  }
		]
	  }
	}
  }`

const SearchProjectsByHashtagsQuery = `{
	"query": {
	  "bool": {
		"must": [
		  {
			"nested": {
			  "path": "hashtags",
			  "query": {
				"terms": {
				  "hashtags.hashtagName": %s
				}
			  }
			}
		  }
		]
	  }
	}
  }`

const FullTextSearchProjectsQuery = `{
	"query": {
	  "multi_match": {
		"query": "%s",
		"fields": ["slug^2", "description"],
		"fuzziness": "AUTO"
	  }
	}
  }`
