import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SportEpisodesComponent } from './sport-episodes.component';

describe('SportEpisodesComponent', () => {
  let component: SportEpisodesComponent;
  let fixture: ComponentFixture<SportEpisodesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SportEpisodesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SportEpisodesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
