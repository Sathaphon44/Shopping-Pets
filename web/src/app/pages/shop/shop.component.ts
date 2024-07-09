import { Component } from '@angular/core';
import { ApiService } from 'src/app/services/api/api.service';
import { DataSharingService } from 'src/app/services/data-sharing/data-sharing.service';

@Component({
    selector: 'app-shop',
    templateUrl: './shop.component.html',
    styleUrls: ['./shop.component.css']
})
export class ShopComponent {

    list_pets: any = [];
    loading: boolean = true;
    list_order: any = []
    isVisibleSelectOrder = false;
    isVisibleDetailPet = false;
    select_pet: any;
    detail_pet: any;
    cartDetail: any;
    closeButton: boolean = false
    close_button: any = {
        confirmBtn: false
    }
    token: string | null = localStorage.getItem("token")

    constructor(
        private http: ApiService,
        private sharing: DataSharingService
    ) { }

    // start http services
    private async FetchDataUseInShopPage() {
        var cart_data: any;
        try {
            if (this.token != null) {
                var shop_data: any = await this.http.ShopServices().toPromise()
                this.sharing.dataPetsInCarts$.subscribe((data: any) => cart_data = {data})
                if (cart_data.data === null) {
                    cart_data = await this.http.GetDetailCart(this.token).toPromise()
                }
            }
            return { shop_data, cart_data }
        } catch (error) {
            return {
                shop_data: undefined,
                cart_data: undefined
            }
        }
    }

    async SetDefaultValueInShopPage() {
        const list_data = await this.FetchDataUseInShopPage()
        console.log(list_data)
        if (list_data.cart_data.data != null) {
            console.log("processing....")
            this.cartDetail = list_data.cart_data.data
            this.list_pets = list_data.shop_data.map((val: any) => {
                const pets_match = this.cartDetail.pets.find((pet: any) => pet.PetId == val.Id)
                return {
                    ...val,
                    amount: pets_match ? pets_match.amount : 0
                }
            });
            this.loading = false
            return
        }
        if (list_data.shop_data != undefined) {
            this.list_pets = list_data.shop_data.map((val: any) => ({ ...val, amount: 0 }));
        }
        this.loading = false
        return
    }


    private async DeleteCart(data: any) {
        try {
            if (data.pet == undefined) {
                if (this.token != null) {
                    try {
                        await this.http.DeleteCart(this.token).toPromise();
                    } catch (error) {
                        console.log(error)
                    }
                    this.cartDetail = undefined;
                }
            }
        } catch (error) {
            console.log("delete cart error: ", error);
        }
    }



    private async UpdateCart(data: any) {
        try {
            this.sharing.sendData({ data: data })
            if (this.token != null) {
                const res: any = await this.http.UpdateDetailCart(this.token, data).toPromise()
                if (res.data.result != null) {
                    this.cartDetail = res.data.result
                    this.sharing.sendData(this.cartDetail)
                    this.list_pets = this.list_pets.map((val: any) => {
                        const matchPet = res?.data.result.pets.find((pet: any) => pet.PetId == val.Id)
                        return {
                            ...val,
                            amount: matchPet ? matchPet.amount : 0
                        }
                    });
                }
            }
        } catch (error) {
            console.log("shop page error: ", error)
        }
    }
    // end http services


    selectPet(index: number, pet: any) {
        this.closeButton = false
        this.isVisibleSelectOrder = true;
        this.select_pet = { ...pet, index };
    }


    handleButtonControl(intent: string) {
        const object = this.select_pet;
        if (object?.amount > 0 && intent === "del") {
            object.amount = object.amount - 1
        } else if (intent === "add") {
            object.amount = object.amount + 1
        }

        this.select_pet = { ...object, total_price: object?.amount * object?.price }
        this.isVisibleSelectOrder = true
    };

    showModalDetailPet(object: any) {
        this.detail_pet = object;
        this.isVisibleDetailPet = true;
    };

    closeShowDetailPet() {
        this.isVisibleDetailPet = false
        this.detail_pet = {};
    };


    async confirmOrder() {
        this.closeButton = true
        const { index, amount } = this.select_pet;
        this.list_pets[index].amount = amount
        const pet: any = this.list_pets[index]
        var newData: any[] = []
        if (this.cartDetail === undefined) {
            newData.push({
                PetId: pet.Id, 
                name: pet.name,
                details: pet.details,
                price: pet.price,
                age: pet.age,
                amount: pet.amount, 
                category: "chihuahua1" ,
            })
        } else {
            const indexCart = this.cartDetail.pets.findIndex((val: any) => val.PetId === pet.Id);
            if (indexCart != -1) {
                if (pet.amount != 0) {
                    this.cartDetail.pets[indexCart].amount = pet.amount
                    newData = this.cartDetail.pets
                } else {
                    this.cartDetail.pets.filter((val: any) => {
                        if (val.PetId !== pet.Id) {
                            newData.push(val)
                        }
                    })
                }
            } else {
                this.cartDetail.pets.filter((val: any) => {
                    if (val.PetId !== pet.Id) {
                        newData.push(val)
                    }
                })
                newData.push({ 
                    PetId: pet.Id, 
                    name: pet.name,
                    details: pet.details,
                    price: pet.price,
                    age: pet.age,
                    amount: pet.amount, 
                    category: "chihuahua1" ,
                })
            }
        }

        const data: any = {
            pets: newData,
            timestamp: new Date().toTimeString()
        }
        if (data.pets.length === 0) {
            await this.DeleteCart(data)
        } else {
            await this.UpdateCart(data)
        }
        setTimeout(() => {
            this.closeButton = false
            this.isVisibleSelectOrder = false;
        }, 2000);
    };

    cancelOrder() {
        this.isVisibleSelectOrder = false;
    };


    ngOnInit() {
        this.SetDefaultValueInShopPage()
    };



}
