package services

import (
	"errors"
	"fmt"

	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/types"
	"gorm.io/gorm"
)

type PayService struct {
	// Dependency Injection for DAOs
	OrderDao    *dao.OrderDao
	UserDao     *dao.UserDao
	ProductDao  *dao.ProductDao
}

func (s *PayService) ProcessPayment(ctx context.Context, req *PaymentRequest) (interface{}, error) {
    

    err = dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error {
        // 获取订单和支付金额
        order, err := dao.NewOrderDaoByDB(tx).GetOrderById(req.OrderId, user.Id)
        if err != nil {
            return err
        }

        totalAmount := order.Money * float64(order.Num)

        // 解密用户余额
        userBalance, err := user.DecryptBalance(req.Key)
        if err != nil || userBalance < totalAmount {
            return errors.New("余额不足")
        }

        // 扣除用户余额并更新
        userBalance -= totalAmount
        user.Money, err = user.EncryptBalance(req.Key, userBalance)
        if err != nil {
            return err
        }

        // 更新用户余额
        err = dao.NewUserDaoByDB(tx).UpdateUserById(user.Id, user)
        if err != nil {
            return err
        }

        // 更新商家余额
        boss, err := dao.NewUserDaoByDB(tx).GetUserById(req.BossID)
        if err != nil {
            return err
        }

        bossBalance, err := boss.DecryptBalance(req.Key)
        if err != nil {
            return err
        }

        bossBalance += totalAmount
        boss.Money, err = boss.EncryptBalance(req.Key, bossBalance)
        if err != nil {
            return err
        }

        err = dao.NewUserDaoByDB(tx).UpdateUserById(req.BossID, boss)
        if err != nil {
            return err
        }

        // 更新商品库存
        product, err := dao.NewProductDaoByDB(tx).GetProductById(req.ProductID)
        if err != nil {
            return err
        }

        product.Num -= order.Num
        err = dao.NewProductDaoByDB(tx).UpdateProduct(req.ProductID, product)
        if err != nil {
            return err
        }

        // 更新订单状态
        order.Type = consts.OrderTypePaid
        err = dao.NewOrderDaoByDB(tx).UpdateOrderById(req.OrderId, user.Id, order)
        if err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    return nil, nil
}
