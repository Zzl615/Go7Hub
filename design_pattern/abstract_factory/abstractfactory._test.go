/**
 * @Author: Noaghzil
 * @Date:   2023-02-11 21:51:57
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-11 21:58:44
 */
package abstractfactory

import (
	"testing"
)

func SaveOrder(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestRDBDAO(t *testing.T) {
	var factory DAOFactory = &RDBDAOFactory{}
	SaveOrder(factory)
}

func TestXMLDAO(t *testing.T) {
	var factory DAOFactory = &XMLDAOFactory{}
	SaveOrder(factory)
}
