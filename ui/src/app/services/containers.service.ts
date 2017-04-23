import {Inject, Injectable} from '@angular/core';
import {Http, Response} from "@angular/http";
import {Observable} from "rxjs/Observable";

@Injectable()
export class ContainersService {

  baseUrl: string;

  constructor(@Inject(Http) private http: Http) {

    this.baseUrl = "/api/containers";
  }

  public all(): Observable<any> {

    return this.http.get(`${this.baseUrl}`)
      .map(r => r.json());
  }
}
