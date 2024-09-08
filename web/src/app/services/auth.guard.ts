import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, GuardResult, MaybeAsync, Router, RouterStateSnapshot } from '@angular/router';
import { UserService } from './user.service';
import { map, Observable, take } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class AuthGuard implements CanActivate {

    constructor(
        private loginService: UserService,
        private router: Router,
    ) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
        return this.loginService.getIsAuthenticated().pipe(
            take(1), // 只取第一个值，避免订阅长时间存在 
            map(isAuthenticated => {
                if (!isAuthenticated) {
                    // 传递参数，以便在用户登录后可以恢复到他们尝试访问的页面。
                    this.router.navigate(['/login'], { queryParams: { returnUrl: state.url } });
                    return false
                }
                return true;
            })
        )
    }
}
