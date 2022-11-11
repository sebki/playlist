package database

var schema string = `
	title: string @index(exact) .
	description: string .
	bggId: string @index(exact) .
	bggType: [string] .
	thumbnail: string .
	image: string .
	links: [Link] .
	minage: int .
	minplayer: int .
	maxplayer: int .
	minplaytime: int .
	maxplaytime: int .
	linktype: uid .
	type Game {
		title
		description
		bggId
		bggType
		thumbnail
		image
		links
		minage
		minplayer
		maxplayer
		minplaytime
		maxplaytime
	}
	type Link {
		linktype
		bggId
		value
		inbound
	}
`

var linkSchema string = `
	
	bggId: string @index(exact) .
	value: string @index(exact) .
	inbound: bool .
	
`
