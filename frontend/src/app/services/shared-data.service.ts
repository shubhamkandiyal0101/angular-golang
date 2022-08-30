import { Injectable } from '@angular/core';
import { BehaviorSubject } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class SharedDataService {

  showHeader = new BehaviorSubject(true);

  constructor() { }
}
