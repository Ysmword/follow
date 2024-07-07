import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class StorageServiceService {

  constructor() { }

  setWithExpiry(key: string, value: any, ttl: number) {
    const now = new Date();
    const item = {
      value: value,
      expiry: now.getTime() + ttl,
    };
    localStorage.setItem(key, JSON.stringify(item));
  }

  get(key: string) {
    const itemStr = localStorage.getItem(key);
    if (!itemStr) {
      return null;
    }
    const item = JSON.parse(itemStr);
    const now = new Date();
    if (now.getTime() > item.expiry) {
      this.remove(key);
      return null;
    }
    return item.value;
  }
 
  remove(key: string) {
    localStorage.removeItem(key);
  }
 
  clear() {
    localStorage.clear();
  }
}
