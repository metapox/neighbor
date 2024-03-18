package writer

func main() {
	config := Config{
		parallel: 10,
	}

	writer := PostgresWriter{
		config: config,
	}

	writer.Connect()
}