package email

func (c *Client) SendWelcomeEmail(to, firstName string) error {
	data := map[string]string {
		"userFirstName": firstName,
	}

	return c.SendEmail(
		to,
		"Welcome to boilerplate!",
		TemplateWelcome,
		data,
	)
}