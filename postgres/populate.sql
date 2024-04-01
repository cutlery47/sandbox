INSERT INTO keywords
	SELECT word, catcode, catdesc FROM pg_get_keywords();


