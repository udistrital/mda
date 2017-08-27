import {Injectable} from '@angular/core';
import {Headers, Http} from '@angular/http';
import {[Entity]} from '../models/[entity]';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class [Entity]Service {

  private [entity-plural]URL = 'http://localhost:8081/v1/[entity]';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  get[EntityPlural](): Promise<[Entity][]> {
    return this.http.get(this.[entity-plural]URL)
      .toPromise()
      .then(response => response.json() as [Entity][])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  get[Entity](id: string): Promise<[Entity]> {
    const url = `${this.[entity-plural]URL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as [Entity])
      .catch(this.handleError);
  }


  update([entity]: [Entity]): Promise<[Entity]> {
    const url = `${this.[entity-plural]URL}/${[entity].id}`;
    return this.http
      .put(url, JSON.stringify([entity-plural]), {headers: this.headers})
      .toPromise()
      .then(() => [entity-plural])
      .catch(this.handleError);
  }


  create([entity]: [Entity]): Promise<[Entity]> {
    return this.http
      .post(this.[entity-plural]URL, JSON.stringify([entity]), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as [Entity])
      .catch(this.handleError);
  }

  delete(id: string): Promise<void> {
    const url = `${this.[entity-plural]URL}/${id}`;
    return this.http.delete(url, {headers: this.headers})
      .toPromise()
      .then(() => null)
      .catch(this.handleError);
  }

}
