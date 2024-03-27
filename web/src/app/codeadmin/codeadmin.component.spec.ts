import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CodeadminComponent } from './codeadmin.component';

describe('CodeadminComponent', () => {
  let component: CodeadminComponent;
  let fixture: ComponentFixture<CodeadminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CodeadminComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CodeadminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
