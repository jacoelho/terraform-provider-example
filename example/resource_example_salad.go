package example

import "github.com/hashicorp/terraform/helper/schema"

func resourceSalad() *schema.Resource {
	return &schema.Resource{
		Create: resourceSaladCreate,
		Read:   resourceSaladRead,
		Update: resourceSaladUpdate,
		Delete: resourceSaladDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Salad Name",
			},
			"size": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "Salad Size",
			},
		},
	}
}

func resourceSaladCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*struct{})

	name := d.Get("name").(string)

	if size, ok := d.GetOk("size"); ok {
		size.(int)
	}

	client.Create(name, size)

	return nil
}

func resourceSaladRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*struct{})

	name := d.Get("name").(string)

	salad, err := client.GetByName(name)
	if err != nil {
		return err
	}

	if err := d.Set("size", salad.Size); err != nil {
		return err
	}

	return nil
}

func resourceSaladUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*struct{})

	name := d.Get("name").(string)
	size := d.Get("size").(int)

	salad, err := client.Update(name, size)
	if err != nil {
		return err
	}

	if err := d.Set("size", salad.Size); err != nil {
		return err
	}

	return nil
}

func resourceSaladDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*struct{})

	name := d.Get("name").(string)

	return client.Delete(name)
}
