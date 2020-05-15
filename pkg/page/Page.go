package page

type Page struct {
	BaseUrl       string `json:"base_url"`
	Size          int    `json:"size"`
	CurrentPage   int    `json:"current_page"`
	TotalPages    int    `json:"total_pages"`
	TotalElements int    `json:"total_elements"`
	RenderPages   []int  `json:"rander_pages"`
	HasNext       bool   `json:"has_next"`
	HasPre        bool   `json:"has_pre"`
}

func NewPage(size, currentPage, totalElements, renderPageNum int, baseUrl string) *Page {
	if renderPageNum%2 != 0 {
		renderPageNum += 1
	}

	totalPages := 0
	if totalElements%size == 0 {
		totalPages = totalElements / size
	} else {
		totalPages = (totalElements / size) + 1
	}

	return &Page{
		BaseUrl:       baseUrl,
		Size:          size,
		CurrentPage:   currentPage,
		TotalPages:    totalPages,
		TotalElements: totalElements,
		RenderPages:   computePageItems(currentPage, totalPages, renderPageNum),
		HasNext:       currentPage < totalPages,
		HasPre:        currentPage > 1,
	}
}

func computePageItems(currentPage, totalPages, renderPageNum int) []int {
	compute := func(min, max int) []int {
		var items []int
		for ; min <= max; min++ {
			items = append(items, min)
		}
		return items
	}

	mid := renderPageNum / 2
	min := 1
	max := 0

	if currentPage == 1 {
		if totalPages-mid > 0 {
			max = currentPage + mid
		} else {
			max = totalPages
		}
	} else {
		if x := currentPage - mid; x > 0 {
			min = currentPage - mid
		} else if x == 0 {
			min = 1
		} else {
			min = currentPage + (currentPage - mid)
		}

		if currentPage+mid > totalPages {
			max = totalPages
		} else {
			max = currentPage + mid
		}
	}

	return compute(min, max)
}
