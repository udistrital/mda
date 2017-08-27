import { Component, OnInit } from '@angular/core';
import { [Entity] } from '../../models/[entity-lower]';
import { [Entity]Service } from '../../services/[entity-lower].service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-[entity-lower]-new',
  templateUrl: './[entity-lower]-new.component.html',
  styleUrls: ['./[entity-lower]-new.component.css']
})
export class [Entity]NewComponent implements OnInit {

  [entity-lower]: [Entity];
  display = false;
  constructor(private [entity-lower]Service: [Entity]Service, private location: Location) { }

  ngOnInit() {
    this.[entity-lower] = new [Entity]();
  }

  guardar([entity-lower]: [Entity]): void {

    this.[entity-lower]Service.create([entity-lower]);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}
