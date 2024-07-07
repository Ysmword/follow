import { Routes } from '@angular/router';
import { CodeComponent } from './code/code.component';
import { CodeadminComponent } from './codeadmin/codeadmin.component';
import { ResultComponent } from './result/result.component';
import { LoginComponent } from './login/login.component';

export const routes: Routes = [
    { path: "code", component: CodeComponent },
    { path: "codeadmin", component: CodeadminComponent },
    { path: "result", component: ResultComponent },
];
