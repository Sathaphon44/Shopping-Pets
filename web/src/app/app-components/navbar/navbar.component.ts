import { Component, Input } from '@angular/core';
import { Router } from '@angular/router';
import { NzModalService } from 'ng-zorro-antd/modal';
import { DataSharingService } from 'src/app/services/data-sharing/data-sharing.service';

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {

    @Input() dataUser: any | null;
    amountPetsInCart: number = 0


    constructor(
        private modal: NzModalService,
        private router: Router,
        private sharing: DataSharingService
    ) { }




    handleLogout() {
        localStorage.removeItem("token")
    }


    showConfirm(): void {
        this.modal.confirm({
            nzCentered: true,
            nzTitle: '<i>คุณต้องการออกจากระบบ ใช่ หรือ ไม่?</i>',
            nzContent: '<b>เมื่อกด ยันยัน จะทำการออกจากระบบ</b>',
            nzOkText: 'ยันยัน',
            nzOnOk: () => {
                localStorage.removeItem("token")
                this.router.navigate(['/']).then(() => window.location.reload())
            },
            nzCancelText: 'ยกเลิก',
            nzOnCancel: () => { }
        });
    }


    ngOnInit(): void {
        this.sharing.dataPetsInCarts$.subscribe(data => {
            if (data != null && data.pets != undefined) {
                this.amountPetsInCart = data?.pets.length
            } else {
                this.amountPetsInCart = 0
            }
        })
    }

}
