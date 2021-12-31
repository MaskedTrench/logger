package logger

// TODO -- Implementation of all level message dispathechers

func filed_format(fields map[string]string) (output []string, err error) {
	for k, v := range fields {
		if k == "" {
			output[0] = "none"
			println(output[0])
			break
		}
		println(k, v)
	}

	return output, nil
}


func (l *Logger) Inform(msg string) {
	filed_format(l.Fields)
	println(msg)
}