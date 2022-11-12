package database

var schema string = `
	title: string @lang @index(fulltext) .
	description: string @lang .
	bggId: string @index(exact) .
	bggType: [string] .
	thumbnail: string .
	image: string .
	links: [uid] .
	minage: int .
	minplayer: int .
	maxplayer: int .
	minplaytime: int .
	maxplaytime: int .
	linktype: string .
	linkvalue: string .
	inbound: bool .
	rank: int .
	game: uid .
	userdescription: string @lang .
	datecreated: dateTime .
	datemodified: dateTime .
	listtype: string .
	length: int .
	games: [uid] .

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
		linkvalue
		inbound
	}
	type ListedGame {
		rank
		userdescription
		game
	}
	type Playlist {
		title
		description
		datecreated
		datemodified
		listtype
		length
		games
	}
`
