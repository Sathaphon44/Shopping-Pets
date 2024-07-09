import { Injectable } from '@angular/core';
import { ApiService } from '../api/api.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CurrentUserService {

  constructor(
    private http: ApiService
  ) { }

  
  VerifyToken(): Observable<any> {
    const token: string | null = localStorage.getItem("token")
    if (token !== null) {
      return this.http.VerifyTokenServices(token)
    } else {
      throw new Error("Token not found.")
    }
  }

  testTest() {
    return true
  }


}
