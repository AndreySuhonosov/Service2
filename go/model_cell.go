/*
 * REST API сервиса интеграции кассетниц ТЛД
 *
 * Интерфейс Прикладного Программирования выдачи ТЛД
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Описание ячейки УВД
type Cell struct {
	// Идентификатор ячейки
	CellID string `json:"Cell_ID,omitempty"`
	// Идентификатор УВД
	DispeneserID string `json:"dispeneserID,omitempty"`
	// Состояние ячейки
	State string `json:"state,omitempty"`
}
