import { Routes } from '@angular/router';
import { CodeComponent } from './code/code.component';
import { CodeadminComponent } from './codeadmin/codeadmin.component';
import { ResultComponent } from './result/result.component';

export const routes: Routes = [
    { path: "", redirectTo: "/code", pathMatch: 'full' },
    { path: "code", component: CodeComponent },
    { path: "codeadmin", component: CodeadminComponent },
    { path: "result", component: ResultComponent },
];
