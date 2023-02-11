/**
 * @Author: Noaghzil
 * @Date:   2023-02-11 21:30:50
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-02-11 21:50:55
 */
package abstractfactory

import "fmt"

//	=== 定义接口 ===
//
// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情纪录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// === OrderDAO的redis实现 ===

type RDBMainDAO struct{}

func (o *RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct{}

func (o *RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

type RDBDAOFactory struct{}

func (o *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (o *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// === OrderDAO的XML实现 ===

type XMLMainDAO struct{}

func (o *XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

type XMLDetailDAO struct{}

func (o *XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

type XMLDAOFactory struct{}

func (o *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (o *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
