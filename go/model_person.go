/*
 * REST API сервиса интеграции кассетниц ТЛД
 *
 * Интерфейс Прикладного Программирования выдачи ТЛД
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Описание работника
type Person struct {
	// Номер пропуска работника
	PassID string `json:"passID,omitempty"`
	// ФИО Работника
	Name string `json:"name,omitempty"`
	// Роль работника если есть: Администратор, инженер
	Role string `json:"role,omitempty"`
}
