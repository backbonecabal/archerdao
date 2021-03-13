/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel } from './base-model';
import { SubTypeFactory } from './sub-type-factory';


export interface IGetPrice_outputs {
    output?: string;
}


export class GetPrice_outputs extends BaseModel implements IGetPrice_outputs  {

    static OUTPUT_FIELD_NAME = 'output';

    /** bytes */
    output: string;

    /**
     * constructor
     * @param values Can be used to set a webapi response or formValues to this newly constructed model
    * @useFormGroupValuesToModel if true use formValues
    */
    constructor(values?: Partial<IGetPrice_outputs>, useFormGroupValuesToModel = false) {
        super();

        if (values) {
            this.setValues(values, useFormGroupValuesToModel);
        }
    }

    /**
     * set the values.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: Partial<IGetPrice_outputs>, useFormGroupValuesToModel = false): void {
        if (values) {
            const rawValues = this.getValuesToUse(values, useFormGroupValuesToModel);
            this.output = this.getValue<string>(rawValues, GetPrice_outputs.OUTPUT_FIELD_NAME);
            // set values in model properties for added formControls
            super.setValuesInAddedPropertiesOfAttachedFormControls(values, useFormGroupValuesToModel);
        }
    }

    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                output: new FormControl(this.output, [Validators.pattern('^(0x)?[a-fA-F0-9]+$'), ]),
            });
        }
        return this._formGroup;
    }

    /**
     * set the FormGroup values to the model values.
    */
    setFormGroupValues() {
        this.$formGroup.controls[GetPrice_outputs.OUTPUT_FIELD_NAME].setValue(this.output);
        // set formValues in added formControls
        super.setFormGroupValuesInAddedFormControls();
    }
}