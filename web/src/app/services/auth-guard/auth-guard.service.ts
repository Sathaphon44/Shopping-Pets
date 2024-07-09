import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
  Router,
  UrlTree,
} from '@angular/router';
import { Observable } from 'rxjs';
import { CurrentUserService } from "../current-user/current-user.service";
@Injectable({
  providedIn: 'root',
})
export class AuthGuard {
  constructor(public cur: CurrentUserService, public router: Router) {}

  async canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Promise<Observable<boolean> | boolean> {

    try {
      const token = localStorage.getItem("token");
      if (token !== null) {
        const res = await this.cur.VerifyToken().toPromise();
        return true;
      } else {
        localStorage.removeItem("token");
        this.router.navigate(['/login']).then(() => window.location.reload());
        return false;
      }
    } catch (error) {
      localStorage.removeItem("token");
      this.router.navigate(['/login']).then(() => window.location.reload());
      return false;
    }
  }
}