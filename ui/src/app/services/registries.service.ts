import {Inject, Injectable} from '@angular/core';
import {Http, Response} from "@angular/http";
import {Observable} from "rxjs/Observable";
import {Registries} from "../../store/data/registries/registry.model";
import NewRegistry = Registries.NewRegistry;

@Injectable()
export class RegistriesService {

  baseUrl: string;

  constructor(@Inject(Http) private http: Http) {

    this.baseUrl = "/api/registries";
  }

  public save(registry: NewRegistry): Observable<any> {

    return this.http
      .post(`${this.baseUrl}`, registry);
  }

  public all(): Observable<any> {

    return this.http
      .get(`${this.baseUrl}`)
      .map(r => r.json());
  }
}
