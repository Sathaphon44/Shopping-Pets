import { Component } from '@angular/core';
import { CurrentUserService } from './services/current-user/current-user.service';
import { ApiService } from './services/api/api.service';
import { DataSharingService } from './services/data-sharing/data-sharing.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  isCollapsed = false;
  dataUser: any | null;
  token: string | null = localStorage.getItem("token")
  constructor(
    private currentUser: CurrentUserService,
    private http: ApiService,
    private sharing: DataSharingService
  ) {}



  async CheckToken() {
    try {
      this.dataUser = await this.currentUser.VerifyToken().toPromise()
      if (this.token != null) {
        const response: any = await this.http.GetDetailCart(this.token).toPromise()
        this.sharing.sendData(response.data)
      }
    } catch (error) {
      this.dataUser = null
    }
  }


  ngOnInit(): void {
    this.CheckToken()
  }
}
