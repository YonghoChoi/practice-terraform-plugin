package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, // address를 필수로 입력하도록 설정
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)

	// 리소스의 ID를 address 값으로 설정 (해당 리소스 값은 tfstate에 기록된다)
	d.SetId(address)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	// 부분 상태 모드로 활성화
	// 이를 활성화하면 리소스의 일부 변경 시 오류가 발생하면 부분적인 상태를 반환한다.
	d.Partial(true)

	if d.HasChange("address") {
		if err := updateAddress(d, m); err != nil {
			return err
		}

		d.SetPartial("address")
	}

	// 부분 모드를 활성화 한채로 함수를 종료하면 변경된 리소스만 반영된다.
	// 부분 모드를 비활성화하면 모든 필드를 다시 저장하게 된다.
	d.Partial(false)

	return resourceServerRead(d, m)
}

func updateAddress(d *schema.ResourceData, m interface{}) error {
	// 예제이기 때문에 별도 업데이트 동작 없음
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
