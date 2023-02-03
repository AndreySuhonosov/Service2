/*
 * REST API сервиса интеграции кассетниц ТЛД
 *
 * Интерфейс Прикладного Программирования выдачи ТЛД
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Описание данных дозиметра
type Dosimeter struct {
	// Идентификатор дозиметра
	Id string `json:"id,omitempty"`
	// Название дозиметра
	Title string `json:"title,omitempty"`
	// Описание дозиметра
	Description string `json:"description,omitempty"`
	// Идентификатор УВД
	DispenserID string `json:"dispenserID,omitempty"`
	// Описание УВД - Место нахождения
	DispenserDesc string `json:"dispenserDesc,omitempty"`
}
