import { TestBed, inject } from '@angular/core/testing';

import { SummariesService } from './summaries.service';

describe('SummariesService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [SummariesService]
    });
  });

  it('should ...', inject([SummariesService], (service: SummariesService) => {
    expect(service).toBeTruthy();
  }));
});
