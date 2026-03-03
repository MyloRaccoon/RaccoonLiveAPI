package anilist

const GetUserIDQuery = `
query ($name: String!) {
	User(name: $name) {
		id
	}
}
`

const GetLastActivityQuery = `
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

const GetFavoritesAnimeQuery = `
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

const GetFavoritesMangaQuery = `
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

const GetFavoritesCharactersQuery = `
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

const GetFavoritesStaffQuery = `
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

const GetFavoritesStudiosQuery = `
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