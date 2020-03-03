Feature: create new versions
  In order to publish a sofware release
  As an automated proces within a pipeline
  we need to be able to create a new version

  Scenario: New project
    Given we have no current version
    Then we should have a new version of "0.0.1"

  Scenario Outline: Release
    Given we have a current version of <current>$
    When we release <level>$
    Then we should have returned a new version of <new>$
    
    Examples:
      | current    | level   | new        |
      | "0.0.1"    | 2       | "0.0.2"    |
      | "0.0.2"    | 1       | "0.1.0"    |
      | "0.1.0"    | 0       | "1.0.0"    | 
      | "2.0.0"    | 0       | "3.0.0"    |   
      | "2.0.0"    | 2       | "2.0.1"    |   
      | "2.0.1"    | 1       | "2.1.0"    |   
      | "4.30.44"  | 0       | "5.0.0"    |   
      | "2.222.45" | 2       | "2.222.46" |   
      | "2.0.0"    | 0       | "3.0.0"    |   



