package managertemplates

type HomePage struct {
	ListTermsURL string
}

func (page HomePage) PageLayout() PageLayout {
	return PageLayout{
		PageTitle: "Home",
	}
}
