import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataSharingService {
  private dataPetsInCarts = new BehaviorSubject<any>(null);
  dataPetsInCarts$ = this.dataPetsInCarts.asObservable();

  constructor() { }

  sendData(data: any) {
    this.dataPetsInCarts.next(data)
  }

}
