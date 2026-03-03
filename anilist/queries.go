package anilist

const getUserIDQuery = `
query ($name: String!) {
	User(name: $name) {
		id
	}
}
`

const getLastActivityQuery = `
query ($userId: Int) {
  Page(page: 1, perPage: 2) {
	activities(userId: $userId, sort: ID_DESC, type: MEDIA_LIST) {
	  ... on ListActivity {
	    id
		status
		progress
		media {
		  title { romaji }
		  siteUrl
		}
	  }
	}
  }
}
`

const getFavoritesAnimeQuery = `
query GetUserFavorites($username: String!) {
	User(name: $username) {
		name
		favourites {
			anime {
				nodes {
					id
					title {
						romaji
						english
						native
					}
					coverImage {
						large
					}
					averageScore
					genres
					siteUrl
				}
			}
		}
	}
}
`

const getFavoritesMangaQuery = `
query GetUserFavorites($username: String!) {
	User(name: $username) {
		name
		favourites {
			manga {
				nodes {
					id
					title {
						romaji
						english
						native
					}
					coverImage {
						large
					}
					averageScore
					genres
					siteUrl
				}
			}
		}
	}
}
`

const getFavoritesCharactersQuery = `
query GetUserFavorites($username: String!) {
	User(name: $username) {
		name
		favourites {
			characters {
        nodes {
          id
          name {
            full
            native
          }
          description
          gender
          dateOfBirth {
          	year
          	month
          	day
          }
          age
          bloodType
          image {
            large
          }
          siteUrl
          media {
          	nodes {
          		id
          		title {
          			romaji
          		}
          		siteUrl
          	}
          }
        }
      }
		}
	}
}
`

const getFavoritesStaffQuery = `
query GetUserFavorites($username: String!) {
	User(name: $username) {
		name
		favourites {
			staff {
        nodes {
          id
          name {
            full
            native
          }
          description
          siteUrl
        }
      }
		}
	}
}
`

const getFavoritesStudiosQuery = `
query GetUserFavorites($username: String!) {
	User(name: $username) {
		name
		favourites {
			studios {
        nodes {
          id
          name
          siteUrl
        }
      }
		}
	}
}
`