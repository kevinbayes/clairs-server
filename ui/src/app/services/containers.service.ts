import {Inject, Injectable} from '@angular/core';
import {Http, Response} from "@angular/http";
import {Observable} from "rxjs/Rx";
import {Containers} from "../../store/data/containers/container.model";
import NewContainer = Containers.NewContainer;

@Injectable()
export class ContainersService {

  baseUrl: string;

  constructor(@Inject(Http) private http: Http) {

    this.baseUrl = "/api/containers";
  }

  public save(container: NewContainer): Observable<any> {

    return this.http.post(`${this.baseUrl}`, container);
  }

  public all(): Observable<any> {

    return this.http.get(`${this.baseUrl}`)
      .map(r => r.json());
  }

  public get(id: number): Observable<any> {

    return this.http.get(`${this.baseUrl}/${id}`)
      .map(r => r.json());
  }
}
