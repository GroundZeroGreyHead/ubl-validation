package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go-basex-validator <xml-file>")
        os.Exit(1)
    }

    xmlFilePath := os.Args[1]
    schematronPath := os.Args[2]

    // Define paths
    // schematronPath := "/basex/11.1/validation-artifacts/en16931-ubl-1.3.12/schematron/EN16931-UBL-validation.sch"
		// baseXDir := "/basex/11.1"
    // schematronPath = filepath.Join(baseXDir, schematronPath)
   // Check if the Schematron file exists
	 	// Check if the XML file exists
		if _, err := os.Stat(xmlFilePath); os.IsNotExist(err) {
				fmt.Printf("XML file does not exist: %s\n", xmlFilePath)
				os.Exit(1)
		}
		if _, err := os.Stat(schematronPath); os.IsNotExist(err) {
				fmt.Printf("Schematron file does not exist: %s\n", schematronPath)
				os.Exit(1)
		}

    // Step 1: Validate XML against Schematron
    fmt.Println("Validating XML against Schematron...")
    // schematronValidationScript := `
    //     xquery version "3.1";
    //     import module namespace sch = "http://www.schematron.com/validator";
    //     declare variable $xml as document-node() external;
    //     declare variable $sch as document-node() external;
    //     try {
    //         sch:validate($xml, $sch),
    //         "Schematron Validation Passed"
    //     } catch * {
    //         "Schematron Validation Failed: " || $err:description
    //     }
    // `

    // cmd := exec.Command(
		// 	"basex", // The BaseX command-line tool executable
		// 	"-i", xmlFilePath, // Specifies the input XML file to be processed
		// 	"-b", fmt.Sprintf("xml=%s", xmlFilePath), // Binds the XML file to the 'xml' variable in the query
		// 	"-b", fmt.Sprintf("sch=%s", schematronPath), // Binds the Schematron file to the 'sch' variable in the query
		// 	schematronValidationScript, // The XQuery script that performs the Schematron validation
		// )
    // output, err := cmd.CombinedOutput()
    // if err != nil {
				
    //     fmt.Printf("Schematron Validation Error: %v\n", err)
    //     fmt.Printf("Output: %s\n", output)
    //     os.Exit(1)
    // }
    // fmt.Printf("Schematron Validation Result:\n%s\n", output)
		  // Build the BaseX command with arguments
			cmd := exec.Command("basex", "-rp", 
			fmt.Sprintf("validate document('%s') with '%s'", xmlFilePath, schematronPath))
	
		// Capture the standard output of the command
		var out bytes.Buffer
		cmd.Stdout = &out
	
		// Run the command
		err := cmd.Run()
		if err != nil {
			fmt.Println(string(out.Bytes()))
			fmt.Println("Error running BaseX command:", err)
			return
		}
	
		// Print the validation output
		fmt.Println(string(out.Bytes()))
}
