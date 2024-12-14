package web3

var InsertNetworkQuery = `
	INSERT INTO network_table
		(network_name, network_url)
	VALUES
		(?, ?)
`

var SelectNetworkList = `
	SELECT network_seq, network_name, network_url
	FROM network_table
	WHERE network_status = 1
	ORDER BY created_date DESC
`

var SelectNetworkByNetworkNameQuery = `
	SELECT network_seq, network_name, network_url
	FROM network_table
	WHERE network_status = 1 AND
		network_name = ?
`