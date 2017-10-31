import {Inject, Injectable} from '@angular/core';
import {Http} from "@angular/http";
import {Observable} from "rxjs/Rx";

@Injectable()
export class SummariesService {

  baseUrl: string;

  constructor(@Inject(Http) private http: Http) {

    this.baseUrl = "/api";
  }

  public overall(): Observable<any> {

    return this.http
      .get(`${this.baseUrl}/summary`)
      .map(r => r.json());
  }

}

