import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, throwError } from 'rxjs';
import { environment } from 'src/environments/environment';


@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(
    private http: HttpClient
  ) { }


  LoginServices(email: string, password: string) {
    const data: any = {
      email: email,
      password: password
    }
    return this.http.post(environment.apiUrl + "/users/login", data).pipe(
      catchError(this.handleError)
    )
  }

  RegisterService(username: string, email: string, password: string) {
    const data: any = {
      username: username,
      email: email,
      password: password
    }
    return this.http.post(environment.apiUrl + "/users/register", data).pipe(
      catchError(this.handleError)
    )
  }

  VerifyTokenServices(token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.get(environment.apiUrl + "/users/session", { headers }).pipe(
      catchError(this.handleError)
    )
  }



  ShopServices() {
    return this.http.get(environment.apiUrl + "/pets/showpetsall").pipe(
      catchError(this.handleError)
    )
  }


  GetDetailCart(token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.get(environment.apiUrl + "/carts/find", { headers })
  }

  UpdateDetailCart(token:string, data: any){
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.post(environment.apiUrl + "/carts/update", data, { headers })
  }

  DeleteCart(token:string){
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.delete(environment.apiUrl + "/carts/delete", { headers })
  }





  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }

}
