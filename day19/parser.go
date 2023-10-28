
  package main

  import (
    "errors"
    "fmt"
    "os"
    "strings"
  )

  func main() {
    content, err := os.ReadFile("./input.txt")
    if err != nil {
      panic(err)
    }

    splitInput := strings.Split(string(content), "\n\n")
    toTry := splitInput[1]
    counter := 0

    toTrySplit := strings.Split(toTry, "\n")
    
    for _, v := range toTrySplit[0: len(toTrySplit) - 1] {
      index := 0
      err := Rule0(v, &index)
      if err == nil && index == len(v) {
        counter++
      }
    }

    fmt.Println(counter)
  }
  
func Rule101(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule64(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule83(message string, index *int) error {
  indexCopy := *index
  err := Rule83First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule83Second(message, index)
  return err
}

func Rule83First(message string, index *int) error {
  err := Rule132(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule83Second(message string, index *int) error {
  err := Rule108(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule61(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule33(message string, index *int) error {
  indexCopy := *index
  err := Rule33First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule33Second(message, index)
  return err
}

func Rule33First(message string, index *int) error {
  err := Rule126(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule33Second(message string, index *int) error {
  err := Rule68(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule80(message string, index *int) error {
  indexCopy := *index
  err := Rule80First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule80Second(message, index)
  return err
}

func Rule80First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule48(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule80Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule91(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule47(message string, index *int) error {
  indexCopy := *index
  err := Rule47First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule47Second(message, index)
  return err
}

func Rule47First(message string, index *int) error {
  err := Rule38(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule47Second(message string, index *int) error {
  err := Rule64(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule40(message string, index *int) error {
  indexCopy := *index
  err := Rule40First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule40Second(message, index)
  return err
}

func Rule40First(message string, index *int) error {
  err := Rule25(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule40Second(message string, index *int) error {
  err := Rule93(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule18(message string, index *int) error {
  indexCopy := *index
  err := Rule18First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule18Second(message, index)
  return err
}

func Rule18First(message string, index *int) error {
  err := Rule38(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule18Second(message string, index *int) error {
  err := Rule108(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule2(message string, index *int) error {
  if string(message[*index]) == "a" {
    *index++
    return nil
  }
  return errors.New("Could not match")
}

func Rule74(message string, index *int) error {
  indexCopy := *index
  err := Rule74First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule74Second(message, index)
  return err
}

func Rule74First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule108(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule74Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule12(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule41(message string, index *int) error {
  indexCopy := *index
  err := Rule41First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule41Second(message, index)
  return err
}

func Rule41First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule17(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule41Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule58(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule56(message string, index *int) error {
  indexCopy := *index
  err := Rule56First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule56Second(message, index)
  return err
}

func Rule56First(message string, index *int) error {
  err := Rule125(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule56Second(message string, index *int) error {
  err := Rule132(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule66(message string, index *int) error {
  indexCopy := *index
  err := Rule66First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule66Second(message, index)
  return err
}

func Rule66First(message string, index *int) error {
  err := Rule12(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule66Second(message string, index *int) error {
  err := Rule27(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule62(message string, index *int) error {
  indexCopy := *index
  err := Rule62First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule62Second(message, index)
  return err
}

func Rule62First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule52(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule62Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule125(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule58(message string, index *int) error {
  indexCopy := *index
  err := Rule58First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule58Second(message, index)
  return err
}

func Rule58First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule128(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule58Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule99(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule89(message string, index *int) error {
  indexCopy := *index
  err := Rule89First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule89Second(message, index)
  return err
}

func Rule89First(message string, index *int) error {
  err := Rule110(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule89Second(message string, index *int) error {
  err := Rule3(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule30(message string, index *int) error {
  indexCopy := *index
  err := Rule30First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule30Second(message, index)
  return err
}

func Rule30First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule85(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule30Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule21(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule20(message string, index *int) error {
  indexCopy := *index
  err := Rule20First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule20Second(message, index)
  return err
}

func Rule20First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule70(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule20Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule111(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule85(message string, index *int) error {
  indexCopy := *index
  err := Rule85First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule85Second(message, index)
  return err
}

func Rule85First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule115(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule85Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule52(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule98(message string, index *int) error {
  indexCopy := *index
  err := Rule98First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule98Second(message, index)
  return err
}

func Rule98First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule46(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule98Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule10(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule12(message string, index *int) error {
  indexCopy := *index
  err := Rule12First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule12Second(message, index)
  return err
}

func Rule12First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule12Second(message string, index *int) error {
  err := Rule43(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule99(message string, index *int) error {
  indexCopy := *index
  err := Rule99First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule99Second(message, index)
  return err
}

func Rule99First(message string, index *int) error {
  err := Rule102(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule99Second(message string, index *int) error {
  err := Rule56(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule127(message string, index *int) error {
  indexCopy := *index
  err := Rule127First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule127Second(message, index)
  return err
}

func Rule127First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule119(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule127Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule100(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule53(message string, index *int) error {
  indexCopy := *index
  err := Rule53First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule53Second(message, index)
  return err
}

func Rule53First(message string, index *int) error {
  err := Rule125(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule53Second(message string, index *int) error {
  err := Rule108(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule15(message string, index *int) error {
  indexCopy := *index
  err := Rule15First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule15Second(message, index)
  return err
}

func Rule15First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule27(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule15Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule108(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule22(message string, index *int) error {
  indexCopy := *index
  err := Rule22First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule22Second(message, index)
  return err
}

func Rule22First(message string, index *int) error {
  err := Rule95(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule22Second(message string, index *int) error {
  err := Rule61(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule109(message string, index *int) error {
  indexCopy := *index
  err := Rule109First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule109Second(message, index)
  return err
}

func Rule109First(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule109Second(message string, index *int) error {
  err := Rule46(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule90(message string, index *int) error {
  indexCopy := *index
  err := Rule90First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule90Second(message, index)
  return err
}

func Rule90First(message string, index *int) error {
  err := Rule76(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule90Second(message string, index *int) error {
  err := Rule121(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule37(message string, index *int) error {
  indexCopy := *index
  err := Rule37First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule37Second(message, index)
  return err
}

func Rule37First(message string, index *int) error {
  err := Rule15(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule37Second(message string, index *int) error {
  err := Rule23(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule27(message string, index *int) error {
  indexCopy := *index
  err := Rule27First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule27Second(message, index)
  return err
}

func Rule27First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule27Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule104(message string, index *int) error {
  indexCopy := *index
  err := Rule104First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule104Second(message, index)
  return err
}

func Rule104First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule46(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule104Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule12(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule32(message string, index *int) error {
  err := Rule64(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule26(message string, index *int) error {
  indexCopy := *index
  err := Rule26First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule26Second(message, index)
  return err
}

func Rule26First(message string, index *int) error {
  err := Rule35(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule26Second(message string, index *int) error {
  err := Rule32(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule123(message string, index *int) error {
  indexCopy := *index
  err := Rule123First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule123Second(message, index)
  return err
}

func Rule123First(message string, index *int) error {
  err := Rule131(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule123Second(message string, index *int) error {
  err := Rule89(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule21(message string, index *int) error {
  indexCopy := *index
  err := Rule21First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule21Second(message, index)
  return err
}

func Rule21First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule132(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule21Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule115(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule95(message string, index *int) error {
  indexCopy := *index
  err := Rule95First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule95Second(message, index)
  return err
}

func Rule95First(message string, index *int) error {
  err := Rule112(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule95Second(message string, index *int) error {
  err := Rule125(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule79(message string, index *int) error {
  indexCopy := *index
  err := Rule79First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule79Second(message, index)
  return err
}

func Rule79First(message string, index *int) error {
  err := Rule88(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule79Second(message string, index *int) error {
  err := Rule60(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule34(message string, index *int) error {
  indexCopy := *index
  err := Rule34First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule34Second(message, index)
  return err
}

func Rule34First(message string, index *int) error {
  err := Rule107(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule34Second(message string, index *int) error {
  err := Rule133(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule6(message string, index *int) error {
  indexCopy := *index
  err := Rule6First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule6Second(message, index)
  return err
}

func Rule6First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule101(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule6Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule110(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule55(message string, index *int) error {
  indexCopy := *index
  err := Rule55First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule55Second(message, index)
  return err
}

func Rule55First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule112(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule55Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule64(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule10(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule115(message string, index *int) error {
  indexCopy := *index
  err := Rule115First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule115Second(message, index)
  return err
}

func Rule115First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule115Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule39(message string, index *int) error {
  indexCopy := *index
  err := Rule39First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule39Second(message, index)
  return err
}

func Rule39First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule61(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule39Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule81(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule42(message string, index *int) error {
  indexCopy := *index
  err := Rule42First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule42Second(message, index)
  return err
}

func Rule42First(message string, index *int) error {
  err := Rule92(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule42Second(message string, index *int) error {
  err := Rule124(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule105(message string, index *int) error {
  indexCopy := *index
  err := Rule105First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule105Second(message, index)
  return err
}

func Rule105First(message string, index *int) error {
  err := Rule38(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule105Second(message string, index *int) error {
  err := Rule10(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule7(message string, index *int) error {
  indexCopy := *index
  err := Rule7First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule7Second(message, index)
  return err
}

func Rule7First(message string, index *int) error {
  err := Rule10(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule7Second(message string, index *int) error {
  err := Rule27(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule106(message string, index *int) error {
  indexCopy := *index
  err := Rule106First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule106Second(message, index)
  return err
}

func Rule106First(message string, index *int) error {
  err := Rule87(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule106Second(message string, index *int) error {
  err := Rule1(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule75(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule46(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule28(message string, index *int) error {
  indexCopy := *index
  err := Rule28First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule28Second(message, index)
  return err
}

func Rule28First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule71(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule28Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule16(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule120(message string, index *int) error {
  indexCopy := *index
  err := Rule120First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule120Second(message, index)
  return err
}

func Rule120First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule27(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule120Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule112(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule100(message string, index *int) error {
  indexCopy := *index
  err := Rule100First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule100Second(message, index)
  return err
}

func Rule100First(message string, index *int) error {
  err := Rule69(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule100Second(message string, index *int) error {
  err := Rule109(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule102(message string, index *int) error {
  indexCopy := *index
  err := Rule102First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule102Second(message, index)
  return err
}

func Rule102First(message string, index *int) error {
  err := Rule46(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule102Second(message string, index *int) error {
  err := Rule132(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule5(message string, index *int) error {
  indexCopy := *index
  err := Rule5First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule5Second(message, index)
  return err
}

func Rule5First(message string, index *int) error {
  err := Rule39(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule5Second(message string, index *int) error {
  err := Rule51(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule14(message string, index *int) error {
  indexCopy := *index
  err := Rule14First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule14Second(message, index)
  return err
}

func Rule14First(message string, index *int) error {
  err := Rule57(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule14Second(message string, index *int) error {
  err := Rule96(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule8(message string, index *int) error {
  err := Rule42(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule25(message string, index *int) error {
  indexCopy := *index
  err := Rule25First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule25Second(message, index)
  return err
}

func Rule25First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule59(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule25Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule67(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule3(message string, index *int) error {
  indexCopy := *index
  err := Rule3First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule3Second(message, index)
  return err
}

func Rule3First(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule3Second(message string, index *int) error {
  err := Rule52(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule59(message string, index *int) error {
  indexCopy := *index
  err := Rule59First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule59Second(message, index)
  return err
}

func Rule59First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule74(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule59Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule83(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule87(message string, index *int) error {
  indexCopy := *index
  err := Rule87First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule87Second(message, index)
  return err
}

func Rule87First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule72(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule87Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule78(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule130(message string, index *int) error {
  indexCopy := *index
  err := Rule130First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule130Second(message, index)
  return err
}

func Rule130First(message string, index *int) error {
  err := Rule111(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule130Second(message string, index *int) error {
  err := Rule117(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule112(message string, index *int) error {
  indexCopy := *index
  err := Rule112First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule112Second(message, index)
  return err
}

func Rule112First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule43(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule112Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule116(message string, index *int) error {
  indexCopy := *index
  err := Rule116First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule116Second(message, index)
  return err
}

func Rule116First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule38(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule116Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule108(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule17(message string, index *int) error {
  indexCopy := *index
  err := Rule17First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule17Second(message, index)
  return err
}

func Rule17First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule122(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule17Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule29(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule124(message string, index *int) error {
  indexCopy := *index
  err := Rule124First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule124Second(message, index)
  return err
}

func Rule124First(message string, index *int) error {
  err := Rule41(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule124Second(message string, index *int) error {
  err := Rule40(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule11(message string, index *int) error {
  err := Rule42(message, index)
  if err != nil {
    return err
  }

  err = Rule31(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule38(message string, index *int) error {
  err := Rule43(message, index)
  if err != nil {
    return err
  }

  err = Rule43(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule44(message string, index *int) error {
  indexCopy := *index
  err := Rule44First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule44Second(message, index)
  return err
}

func Rule44First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule123(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule44Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule114(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule23(message string, index *int) error {
  err := Rule12(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule29(message string, index *int) error {
  indexCopy := *index
  err := Rule29First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule29Second(message, index)
  return err
}

func Rule29First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule113(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule29Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule129(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule121(message string, index *int) error {
  indexCopy := *index
  err := Rule121First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule121Second(message, index)
  return err
}

func Rule121First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule10(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule121Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule19(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule93(message string, index *int) error {
  indexCopy := *index
  err := Rule93First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule93Second(message, index)
  return err
}

func Rule93First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule26(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule93Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule36(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule13(message string, index *int) error {
  indexCopy := *index
  err := Rule13First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule13Second(message, index)
  return err
}

func Rule13First(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule13Second(message string, index *int) error {
  err := Rule46(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule107(message string, index *int) error {
  indexCopy := *index
  err := Rule107First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule107Second(message, index)
  return err
}

func Rule107First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule132(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule107Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule10(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule54(message string, index *int) error {
  indexCopy := *index
  err := Rule54First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule54Second(message, index)
  return err
}

func Rule54First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule27(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule54Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule12(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule63(message string, index *int) error {
  indexCopy := *index
  err := Rule63First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule63Second(message, index)
  return err
}

func Rule63First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule105(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule63Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule82(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule84(message string, index *int) error {
  indexCopy := *index
  err := Rule84First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule84Second(message, index)
  return err
}

func Rule84First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule37(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule84Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule73(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule52(message string, index *int) error {
  indexCopy := *index
  err := Rule52First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule52Second(message, index)
  return err
}

func Rule52First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule52Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule50(message string, index *int) error {
  indexCopy := *index
  err := Rule50First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule50Second(message, index)
  return err
}

func Rule50First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule111(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule50Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule9(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule111(message string, index *int) error {
  indexCopy := *index
  err := Rule111First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule111Second(message, index)
  return err
}

func Rule111First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule112(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule111Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule27(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule82(message string, index *int) error {
  indexCopy := *index
  err := Rule82First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule82Second(message, index)
  return err
}

func Rule82First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule46(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule82Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule108(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule118(message string, index *int) error {
  if string(message[*index]) == "b" {
    *index++
    return nil
  }
  return errors.New("Could not match")
}

func Rule133(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule103(message string, index *int) error {
  indexCopy := *index
  err := Rule103First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule103Second(message, index)
  return err
}

func Rule103First(message string, index *int) error {
  err := Rule34(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule103Second(message string, index *int) error {
  err := Rule22(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule36(message string, index *int) error {
  indexCopy := *index
  err := Rule36First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule36Second(message, index)
  return err
}

func Rule36First(message string, index *int) error {
  err := Rule66(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule36Second(message string, index *int) error {
  err := Rule65(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule31(message string, index *int) error {
  indexCopy := *index
  err := Rule31First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule31Second(message, index)
  return err
}

func Rule31First(message string, index *int) error {
  err := Rule33(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule31Second(message string, index *int) error {
  err := Rule79(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule69(message string, index *int) error {
  indexCopy := *index
  err := Rule69First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule69Second(message, index)
  return err
}

func Rule69First(message string, index *int) error {
  err := Rule12(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule69Second(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule122(message string, index *int) error {
  indexCopy := *index
  err := Rule122First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule122Second(message, index)
  return err
}

func Rule122First(message string, index *int) error {
  err := Rule54(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule122Second(message string, index *int) error {
  err := Rule66(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule114(message string, index *int) error {
  indexCopy := *index
  err := Rule114First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule114Second(message, index)
  return err
}

func Rule114First(message string, index *int) error {
  err := Rule6(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule114Second(message string, index *int) error {
  err := Rule80(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule92(message string, index *int) error {
  indexCopy := *index
  err := Rule92First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule92Second(message, index)
  return err
}

func Rule92First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule14(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule92Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule44(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule0(message string, index *int) error {
  err := Rule8(message, index)
  if err != nil {
    return err
  }

  err = Rule11(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule117(message string, index *int) error {
  indexCopy := *index
  err := Rule117First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule117Second(message, index)
  return err
}

func Rule117First(message string, index *int) error {
  err := Rule132(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule117Second(message string, index *int) error {
  err := Rule108(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule76(message string, index *int) error {
  indexCopy := *index
  err := Rule76First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule76Second(message, index)
  return err
}

func Rule76First(message string, index *int) error {
  err := Rule64(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule76Second(message string, index *int) error {
  err := Rule27(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule126(message string, index *int) error {
  indexCopy := *index
  err := Rule126First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule126Second(message, index)
  return err
}

func Rule126First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule28(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule126Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule127(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule113(message string, index *int) error {
  indexCopy := *index
  err := Rule113First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule113Second(message, index)
  return err
}

func Rule113First(message string, index *int) error {
  err := Rule125(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule113Second(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule43(message string, index *int) error {
  indexCopy := *index
  err := Rule43First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule43Second(message, index)
  return err
}

func Rule43First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule43Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule46(message string, index *int) error {
  indexCopy := *index
  err := Rule46First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule46Second(message, index)
  return err
}

func Rule46First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule46Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule132(message string, index *int) error {
  indexCopy := *index
  err := Rule132First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule132Second(message, index)
  return err
}

func Rule132First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule132Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule49(message string, index *int) error {
  indexCopy := *index
  err := Rule49First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule49Second(message, index)
  return err
}

func Rule49First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule130(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule49Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule50(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule129(message string, index *int) error {
  indexCopy := *index
  err := Rule129First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule129Second(message, index)
  return err
}

func Rule129First(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule129Second(message string, index *int) error {
  err := Rule27(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule24(message string, index *int) error {
  indexCopy := *index
  err := Rule24First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule24Second(message, index)
  return err
}

func Rule24First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule27(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule24Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule52(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule119(message string, index *int) error {
  indexCopy := *index
  err := Rule119First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule119Second(message, index)
  return err
}

func Rule119First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule47(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule119Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule98(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule91(message string, index *int) error {
  indexCopy := *index
  err := Rule91First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule91Second(message, index)
  return err
}

func Rule91First(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule91Second(message string, index *int) error {
  err := Rule10(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule35(message string, index *int) error {
  indexCopy := *index
  err := Rule35First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule35Second(message, index)
  return err
}

func Rule35First(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule35Second(message string, index *int) error {
  err := Rule46(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule57(message string, index *int) error {
  indexCopy := *index
  err := Rule57First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule57Second(message, index)
  return err
}

func Rule57First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule63(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule57Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule90(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule125(message string, index *int) error {
  indexCopy := *index
  err := Rule125First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule125Second(message, index)
  return err
}

func Rule125First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule125Second(message string, index *int) error {
  err := Rule43(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule108(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule4(message string, index *int) error {
  indexCopy := *index
  err := Rule4First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule4Second(message, index)
  return err
}

func Rule4First(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule4Second(message string, index *int) error {
  err := Rule19(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule94(message string, index *int) error {
  indexCopy := *index
  err := Rule94First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule94Second(message, index)
  return err
}

func Rule94First(message string, index *int) error {
  err := Rule112(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule94Second(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule1(message string, index *int) error {
  indexCopy := *index
  err := Rule1First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule1Second(message, index)
  return err
}

func Rule1First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule55(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule1Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule75(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule9(message string, index *int) error {
  indexCopy := *index
  err := Rule9First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule9Second(message, index)
  return err
}

func Rule9First(message string, index *int) error {
  err := Rule38(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule9Second(message string, index *int) error {
  err := Rule112(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule71(message string, index *int) error {
  indexCopy := *index
  err := Rule71First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule71Second(message, index)
  return err
}

func Rule71First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule18(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule71Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule53(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule60(message string, index *int) error {
  indexCopy := *index
  err := Rule60First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule60Second(message, index)
  return err
}

func Rule60First(message string, index *int) error {
  err := Rule103(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule60Second(message string, index *int) error {
  err := Rule86(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule51(message string, index *int) error {
  indexCopy := *index
  err := Rule51First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule51Second(message, index)
  return err
}

func Rule51First(message string, index *int) error {
  err := Rule116(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule51Second(message string, index *int) error {
  err := Rule24(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule128(message string, index *int) error {
  indexCopy := *index
  err := Rule128First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule128Second(message, index)
  return err
}

func Rule128First(message string, index *int) error {
  err := Rule24(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule128Second(message string, index *int) error {
  err := Rule117(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule45(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule16(message string, index *int) error {
  indexCopy := *index
  err := Rule16First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule16Second(message, index)
  return err
}

func Rule16First(message string, index *int) error {
  err := Rule104(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule16Second(message string, index *int) error {
  err := Rule62(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule70(message string, index *int) error {
  indexCopy := *index
  err := Rule70First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule70Second(message, index)
  return err
}

func Rule70First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule115(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule70Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule45(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule77(message string, index *int) error {
  indexCopy := *index
  err := Rule77First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule77Second(message, index)
  return err
}

func Rule77First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule13(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule77Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule94(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule131(message string, index *int) error {
  indexCopy := *index
  err := Rule131First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule131Second(message, index)
  return err
}

func Rule131First(message string, index *int) error {
  err := Rule61(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule131Second(message string, index *int) error {
  err := Rule107(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule68(message string, index *int) error {
  indexCopy := *index
  err := Rule68First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule68Second(message, index)
  return err
}

func Rule68First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule49(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule68Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule84(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule86(message string, index *int) error {
  indexCopy := *index
  err := Rule86First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule86Second(message, index)
  return err
}

func Rule86First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule97(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule86Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule30(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule78(message string, index *int) error {
  indexCopy := *index
  err := Rule78First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule78Second(message, index)
  return err
}

func Rule78First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule38(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule78Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule12(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule88(message string, index *int) error {
  indexCopy := *index
  err := Rule88First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule88Second(message, index)
  return err
}

func Rule88First(message string, index *int) error {
  err := Rule5(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule88Second(message string, index *int) error {
  err := Rule106(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule81(message string, index *int) error {
  indexCopy := *index
  err := Rule81First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule81Second(message, index)
  return err
}

func Rule81First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule45(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule81Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule52(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule64(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

func Rule19(message string, index *int) error {
  indexCopy := *index
  err := Rule19First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule19Second(message, index)
  return err
}

func Rule19First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule43(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule19Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule110(message string, index *int) error {
  indexCopy := *index
  err := Rule110First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule110Second(message, index)
  return err
}

func Rule110First(message string, index *int) error {
  err := Rule132(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule110Second(message string, index *int) error {
  err := Rule64(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule96(message string, index *int) error {
  indexCopy := *index
  err := Rule96First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule96Second(message, index)
  return err
}

func Rule96First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule77(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule96Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule20(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule67(message string, index *int) error {
  indexCopy := *index
  err := Rule67First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule67Second(message, index)
  return err
}

func Rule67First(message string, index *int) error {
  err := Rule104(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule67Second(message string, index *int) error {
  err := Rule4(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule65(message string, index *int) error {
  indexCopy := *index
  err := Rule65First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule65Second(message, index)
  return err
}

func Rule65First(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule65Second(message string, index *int) error {
  err := Rule38(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule73(message string, index *int) error {
  indexCopy := *index
  err := Rule73First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule73Second(message, index)
  return err
}

func Rule73First(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule110(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule73Second(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule120(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule72(message string, index *int) error {
  indexCopy := *index
  err := Rule72First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule72Second(message, index)
  return err
}

func Rule72First(message string, index *int) error {
  err := Rule64(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule72Second(message string, index *int) error {
  err := Rule115(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule48(message string, index *int) error {
  indexCopy := *index
  err := Rule48First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule48Second(message, index)
  return err
}

func Rule48First(message string, index *int) error {
  err := Rule2(message, index)
  if err != nil {
    return err
  }

  err = Rule46(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule48Second(message string, index *int) error {
  err := Rule118(message, index)
  if err != nil {
    return err
  }

  err = Rule10(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule97(message string, index *int) error {
  indexCopy := *index
  err := Rule97First(message, &indexCopy)

  if err == nil {
    *index = indexCopy
    return nil
  }

  err = Rule97Second(message, index)
  return err
}

func Rule97First(message string, index *int) error {
  err := Rule7(message, index)
  if err != nil {
    return err
  }

  err = Rule118(message, index)
  if err != nil {
    return err
  }

  return nil
}


func Rule97Second(message string, index *int) error {
  err := Rule133(message, index)
  if err != nil {
    return err
  }

  err = Rule2(message, index)
  if err != nil {
    return err
  }

  return nil
}

