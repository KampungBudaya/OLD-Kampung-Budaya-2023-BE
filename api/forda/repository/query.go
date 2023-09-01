package repository

const queryCreateForda = `
	INSERT INTO
		users
	(
		name,
		email,
		phone
	) VALUES (
		?,
		?,
		?
	)
`

const queryCreatePhoto = `
	INSERT INTO
		payment_photos
	(
		link_photo,
		user_id
	) VALUES (
		?,
		?
	)
`

const queryFindForda = `
	SELECT
		users.id,
		users.provider,
		users.provider_id,
		users.name,
		users.email,
		users.password,
		users.phone,
		users.created_at,
		users.updated_at,
		payment_photos.link_photo
	FROM
		users
	LEFT JOIN payment_photos ON users.id = payment_photos.user_id
	%s
`
