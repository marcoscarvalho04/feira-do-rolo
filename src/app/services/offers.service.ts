import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { catchError, delay, Observable, retry } from 'rxjs';
import {Offers} from '../model/offers'

@Injectable({
  providedIn: 'root'
})
export class OffersService {

  private offersApi: string;

  constructor(private http: HttpClient) {
    this.offersApi = "http://localhost:8080/offers"
   }

   public findAll(): Observable<Offers[]> {
   return this.http.get<Offers[]>(this.offersApi)
    .pipe(
      retry(5),
      delay(100)
    );

   }

   public save(offer: Offers): Observable<Offers> {
    return this.http.post<Offers>(this.offersApi, offer)
    .pipe(
      retry(5)
    )

  }

  }
