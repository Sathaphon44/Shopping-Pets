import { Component } from '@angular/core';
import { ApiService } from 'src/app/services/api/api.service';
import { DataSharingService } from 'src/app/services/data-sharing/data-sharing.service';

@Component({
    selector: 'app-cart',
    templateUrl: './cart.component.html',
    styleUrls: ['./cart.component.css']
})
export class CartComponent {

    listDataDetailCart: any | undefined = {}
    loading: boolean = true
    token: string | null = localStorage.getItem("token")
    
    constructor(
        private http: ApiService,
        private sharing: DataSharingService
    ) { }


    private async fetchDataUseInCartPage() {
        try {
            if (this.token !== null) {
                var res: any = await this.http.GetDetailCart(this.token).toPromise()
            }
            return {cart_list: res}
        } catch (error) {
            return {cart_list: undefined}
        }
    }


    private async fetchData() {
        const { cart_list } = await this.fetchDataUseInCartPage()
        if (cart_list.data != null) {
            this.listDataDetailCart = cart_list?.data
            this.loading = false
            return
        }
        this.listDataDetailCart = undefined
        this.loading = false
        return
    }


    public async handleClearCart() {
        this.loading = true
        if (this.token != null) {
            try {
                await this.http.DeleteCart(this.token).toPromise()
            } catch (error) {console.log(error)}
        }
        setTimeout(() => {
            this.listDataDetailCart = undefined;
            this.sharing.sendData(null)
            this.loading = false;
        }, 2000)
    }


    delete_pet(data: any) {
        console.log(data)
    }


    ngOnInit(): void {
        this.fetchData()
    }
}
