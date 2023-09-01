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
		id,
		provider,
		provider_id,
		name,
		email,
		password,
		phone,
		created_at,
		updated_at
	FROM
		users
	%s
`
