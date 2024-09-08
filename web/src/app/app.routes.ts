import { Routes } from '@angular/router';
import { CodeComponent } from './code/code.component';
import { CodeadminComponent } from './codeadmin/codeadmin.component';
import { ResultComponent } from './result/result.component';
import { LoginComponent } from './login/login.component';
import { AuthGuard } from './services/auth.guard';

export const routes: Routes = [
    { path: "code", component: CodeComponent, canActivate: [AuthGuard] },
    { path: "codeadmin", component: CodeadminComponent, canActivate: [AuthGuard] },
    { path: "result", component: ResultComponent, canActivate: [AuthGuard] },
    { path: "login", component: LoginComponent }
];
